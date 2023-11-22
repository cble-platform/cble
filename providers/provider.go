package providers

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/provider"
	"github.com/cble-platform/cble-backend/ent/providercommand"
	"github.com/cble-platform/cble-backend/internal/git"
	"github.com/cble-platform/cble-provider-grpc/pkg/common"
	providerGRPC "github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

func (ps *CBLEServer) downloadProvider(entProvider *ent.Provider) error {
	providerRepoPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "source")
	logrus.WithFields(logrus.Fields{"repo_path": providerRepoPath}).Debugf("Downloading provider %s", entProvider.ID.String())

	// Clone/checkout the provider from git
	if _, err := os.Stat(providerRepoPath); os.IsNotExist(err) {
		logrus.Debugf("Provider does not exist, cloning repo")
		// Provider dir doesn't exist so clone repo
		err := git.CloneProvider(providerRepoPath, entProvider)
		if err != nil {
			return fmt.Errorf("failed to clone provider repo: %v", err)
		}
	} else {
		logrus.Debugf("Provider exists, checking out version")
		// Provider dir exists so just checkout new version
		err := git.CheckoutProvider(providerRepoPath, entProvider)
		if err != nil {
			return fmt.Errorf("failed to checkout provider repo: %v", err)
		}
	}

	providerBinaryPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "provider")
	logrus.WithFields(logrus.Fields{"binary_path": providerBinaryPath}).Debugf("Compiling provider %s", entProvider.ID.String())

	// Build the provider into a binary
	cmd := exec.Command("sh", "-c", fmt.Sprintf("go get ./... && go build -o %s %s", providerBinaryPath, providerRepoPath))
	cmd.Dir = providerRepoPath
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("command exited with status %d. output: %s", exiterr.ExitCode(), cmdOutput)
		}
		return fmt.Errorf("command run error: %v", err)
	}

	return nil
}

// Runs a provider binary. Should be run as a go routine
func (ps *CBLEServer) runProvider(ctx context.Context, entProvider *ent.Provider, shutdown chan bool) {
	providerBinaryPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "provider")

	// Check the provider is compiled
	if _, err := os.Stat(providerBinaryPath); os.IsNotExist(err) {
		// The provider binary has yet to be built
		logrus.Errorf("failed to run provider server: provider has not been compiled yet")
		return
	}

	logrus.Debugf("Executing provider server binary for %s", entProvider.ID.String())

	// Start the binary with the provider ID as argument
	cmd := exec.Command(providerBinaryPath, entProvider.ID.String())
	if err := cmd.Start(); err != nil {
		logrus.Errorf("failed to run provider server: failed to start provider: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			logrus.Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			cmd.Process.Signal(syscall.SIGTERM)
			return
		case <-shutdown:
			logrus.Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			cmd.Process.Signal(syscall.SIGTERM)
			return
		}
	}
}

func (ps *CBLEServer) startProviderConnection(ctx context.Context, shutdown chan bool, providerKey string) {
	registeredProvider, exists := ps.registeredProviders.Load(providerKey)
	if !exists {
		logrus.Errorf("attempted to start provider on non-registered provider (%s)", providerKey)
		return
	}

	logrus.Debugf("starting provider connection to provider %s with socket ID %s", providerKey, registeredProvider.(RegisteredProvider).SocketID)

	providerOpts := &providerGRPC.ProviderClientOptions{
		// TODO: implement TLS for provider connections
		TLS:      false,
		CAFile:   "",
		SocketID: registeredProvider.(RegisteredProvider).SocketID,
	}
	providerConn, err := providerGRPC.Connect(providerOpts)
	if err != nil {
		logrus.Errorf("failed to connect to provider gRPC server (%s): %v", providerKey, err)
		return
	}
	client, err := providerGRPC.NewClient(ctx, providerConn)
	if err != nil {
		logrus.Errorf("failed to create client for provider (%s): %v", providerKey, err)
		return
	}

	// Convert provider ID to UUID for ENT queries
	providerUuid, err := uuid.Parse(providerKey)
	if err != nil {
		logrus.Errorf("failed to parse provider key \"%s\" when starting provider connection: %v", providerKey, err)
		return
	}

	// Provider connection event loop
	for {
		select {
		case <-shutdown:
			logrus.Warnf("Gracefully shutting down provider client %s", providerKey)
			return
		case <-ctx.Done():
			logrus.Warnf("Gracefully shutting down provider client %s", providerKey)
			return
		default:
			// If not cancelling, query ent for all queued commands for this provider
			entCommands, err := ps.entClient.ProviderCommand.Query().Where(
				providercommand.And(
					providercommand.StatusEQ(providercommand.StatusQUEUED),
					providercommand.HasProviderWith(provider.IDEQ(providerUuid)),
				),
			).All(ctx)
			if err != nil {
				logrus.Errorf("failed to query commands for provider \"%s\": %v", providerKey, err)
				continue
			}

			// Run all of the commands in go routines
			for _, entCommand := range entCommands {
				go ps.handleProviderCommand(ctx, client, entCommand)
			}

			// Wait 10 seconds before querying again
			time.Sleep(10 * time.Second)
		}
	}
}

func (ps *CBLEServer) handleProviderCommand(ctx context.Context, client providerGRPC.ProviderClient, entCommand *ent.ProviderCommand) {
	// Set the command start time and mark as in progress
	err := entCommand.Update().SetStatus(providercommand.StatusINPROGRESS).SetStartTime(time.Now()).Exec(ctx)
	if err != nil {
		logrus.Errorf("failed to set command status to in progress and start time to time.Now()")
		// Proceed as this isn't a critical failure
	}

	switch entCommand.CommandType {
	case providercommand.CommandTypeCONFIGURE:
		// Get the provider associated with command
		entProvider, err := entCommand.QueryProvider().Only(ctx)
		if err != nil {
			failCommand(ctx, entCommand, "failed to query provider from command", err)
			return
		}

		// Generate the configuration command
		configureCommand := &providerGRPC.ConfigureRequest{
			Config: entProvider.ConfigBytes,
		}

		// Send the configuration request
		reply, err := client.Configure(ctx, configureCommand)
		if err != nil {
			failCommand(ctx, entCommand, "failed to call provider configure", err)
			return
		}

		// Update the output of the command
		err = entCommand.Update().
			SetStatus(providercommand.StatusSUCCEEDED).
			SetOutput(fmt.Sprintf("RPC status is:\n%s", reply.Status.String())).
			SetEndTime(time.Now()).
			Exec(ctx)
		if err != nil {
			logrus.Errorf("failed to update command state and output")
		}

	case providercommand.CommandTypeDEPLOY:
		// Get the deployment and blueprint associated with command
		entDeployment, err := entCommand.QueryDeployment().Only(ctx)
		if err != nil {
			failCommand(ctx, entCommand, "failed to query deployment from command", err)
			return
		}
		entBlueprint, err := entDeployment.QueryBlueprint().Only(ctx)
		if err != nil {
			failCommand(ctx, entCommand, "failed to query blueprint from deployment", err)
			return
		}

		// Convert maps into protobuf-friendly structs
		templateVarsStruct, err := structpb.NewStruct(entDeployment.TemplateVars)
		if err != nil {
			failCommand(ctx, entCommand, "failed to parse template vars into structpb", err)
			return
		}
		deploymentVarsStruct, err := structpb.NewStruct(entDeployment.DeploymentVars)
		if err != nil {
			failCommand(ctx, entCommand, "failed to parse deployment vars into structpb", err)
			return
		}
		// Deployment state is of type map[string]string and needs to be converted to map[string]interface{}
		deploymentState := make(map[string]interface{}, len(entDeployment.DeploymentState))
		for k, v := range entDeployment.DeploymentState {
			deploymentState[k] = v
		}
		deploymentStateStruct, err := structpb.NewStruct(deploymentState)
		if err != nil {
			failCommand(ctx, entCommand, "failed to parse deployment state into structpb", err)
			return
		}

		// Generate the deployment command
		deployCommand := &providerGRPC.DeployRequest{
			DeploymentId:    entDeployment.ID.String(),
			Blueprint:       entBlueprint.BlueprintTemplate,
			TemplateVars:    templateVarsStruct,
			DeploymentState: deploymentStateStruct,
			DeploymentVars:  deploymentVarsStruct,
		}

		// Send the deploy request
		reply, err := client.Deploy(ctx, deployCommand)
		if err != nil {
			failCommand(ctx, entCommand, "failed to call provider deploy", err)
			return
		}

		// Convert deployment state from map[string]interface{} to map[string]string
		newDeploymentState := make(map[string]string, len(reply.DeploymentState.AsMap()))
		for k, v := range reply.DeploymentState.AsMap() {
			stateVal, ok := v.(string)
			if ok {
				newDeploymentState[k] = stateVal
			} else {
				logrus.Warnf("deployment state value of %v is not string type as expected", v)
			}
		}

		// Update the deployment with the resulting state and variables
		err = entDeployment.Update().
			SetDeploymentState(newDeploymentState).
			SetDeploymentVars(reply.DeploymentVars.AsMap()).
			Exec(ctx)
		if err != nil {
			failCommand(ctx, entCommand, "failed to update deployment state and vars", err)
			return
		}

		var status providercommand.Status
		switch reply.Status {
		case common.RPCStatus_FAILURE:
			status = providercommand.StatusFAILED
		default:
			status = providercommand.StatusSUCCEEDED
		}

		// Update the output of the command
		err = entCommand.Update().
			SetStatus(status).
			SetOutput(fmt.Sprintf("RPC status is:\n%s", reply.Status.String())).
			SetError(fmt.Sprintf("Errors:\n%s", strings.Join(reply.Errors, "\n"))).
			SetEndTime(time.Now()).
			Exec(ctx)
		if err != nil {
			logrus.Errorf("failed to update command state and output")
		}

	case providercommand.CommandTypeDESTROY:
		// Get the deployment and blueprint associated with command
		entDeployment, err := entCommand.QueryDeployment().Only(ctx)
		if err != nil {
			failCommand(ctx, entCommand, "failed to query deployment from command", err)
			return
		}
		entBlueprint, err := entDeployment.QueryBlueprint().Only(ctx)
		if err != nil {
			failCommand(ctx, entCommand, "failed to query blueprint from deployment", err)
			return
		}

		// Convert maps into protobuf-friendly structs
		templateVarsStruct, err := structpb.NewStruct(entDeployment.TemplateVars)
		if err != nil {
			failCommand(ctx, entCommand, "failed to parse template vars into structpb", err)
			return
		}
		deploymentVarsStruct, err := structpb.NewStruct(entDeployment.DeploymentVars)
		if err != nil {
			failCommand(ctx, entCommand, "failed to parse deployment vars into structpb", err)
			return
		}
		// Deployment state is of type map[string]string and needs to be converted to map[string]interface{}
		deploymentState := make(map[string]interface{}, len(entDeployment.DeploymentState))
		for k, v := range entDeployment.DeploymentState {
			deploymentState[k] = v
		}
		deploymentStateStruct, err := structpb.NewStruct(deploymentState)
		if err != nil {
			failCommand(ctx, entCommand, "failed to parse deployment state into structpb", err)
			return
		}

		// Generate the deployment command
		destroyCommand := &providerGRPC.DestroyRequest{
			DeploymentId:    entDeployment.ID.String(),
			Blueprint:       entBlueprint.BlueprintTemplate,
			TemplateVars:    templateVarsStruct,
			DeploymentState: deploymentStateStruct,
			DeploymentVars:  deploymentVarsStruct,
		}

		// Send the destroy request
		reply, err := client.Destroy(ctx, destroyCommand)
		if err != nil {
			failCommand(ctx, entCommand, "failed to call provider destroy", err)
			return
		}

		// Convert deployment state from map[string]interface{} to map[string]string
		newDeploymentState := make(map[string]string, len(reply.DeploymentState.AsMap()))
		for k, v := range reply.DeploymentState.AsMap() {
			stateVal, ok := v.(string)
			if ok {
				newDeploymentState[k] = stateVal
			} else {
				logrus.Warnf("deployment state value of %v is not string type as expected", v)
			}
		}

		// Update the deployment with the resulting state and variables
		err = entDeployment.Update().
			SetDeploymentState(newDeploymentState).
			SetDeploymentVars(reply.DeploymentVars.AsMap()).
			Exec(ctx)
		if err != nil {
			failCommand(ctx, entCommand, "failed to update deployment state and vars", err)
			return
		}

		var status providercommand.Status
		switch reply.Status {
		case common.RPCStatus_FAILURE:
			status = providercommand.StatusFAILED
		default:
			status = providercommand.StatusSUCCEEDED
		}

		// Update the output of the command
		err = entCommand.Update().
			SetStatus(status).
			SetOutput(fmt.Sprintf("RPC status is:\n%s", reply.Status.String())).
			SetError(fmt.Sprintf("Errors:\n%s", strings.Join(reply.Errors, "\n"))).
			SetEndTime(time.Now()).
			Exec(ctx)
		if err != nil {
			logrus.Errorf("failed to update command state and output")
		}
	}
}

func failCommand(ctx context.Context, entCommand *ent.ProviderCommand, message string, err error) {
	updateErr := entCommand.Update().
		SetStatus(providercommand.StatusFAILED).
		SetError(fmt.Sprintf("%s: %v", message, err)).
		SetEndTime(time.Now()).
		Exec(ctx)
	if updateErr != nil {
		logrus.Errorf("failed to update command state and error: %v", updateErr)
	}
}
