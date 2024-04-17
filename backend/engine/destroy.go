package engine

import (
	"context"
	"sync"
	"time"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/ent/resource"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/sirupsen/logrus"
)

// Run in a go routine, destroys the deployment with a provider
func StartDestroy(client *ent.Client, cbleServer *providers.CBLEServer, entDeployment *ent.Deployment) {
	// Create new context for deployment
	ctx := context.Background()

	// Set the deployment as IN_PROGRESS
	entDeployment, err := entDeployment.Update().
		SetState(deployment.StateInProgress).
		Save(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "DESTROY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to update deployment state: %v", err)
		return
	}

	// Set all of the nodes to TO_DESTROY state (which aren't destroyed)
	err = client.DeploymentNode.Update().
		Where(
			deploymentnode.And(
				deploymentnode.HasDeploymentWith(deployment.IDEQ(entDeployment.ID)),
				deploymentnode.StateNEQ(deploymentnode.StateDestroyed), // Only non-destroyed nodes
			),
		).
		SetState(deploymentnode.StateToDestroy).
		Exec(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "DESTROY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to set deployment nodes to TO_DESTROY: %v", err)
	}

	// Query all of the deployment nodes
	entDeploymentNodes, err := entDeployment.QueryDeploymentNodes().All(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "DESTROY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to query deployment nodes: %v", err)
		return
	}

	var wg sync.WaitGroup

	// Spawn destroyRoutine's for all root nodes
	for _, entDeploymentNode := range entDeploymentNodes {
		wg.Add(1)
		go destroyRoutine(ctx, cbleServer, entDeploymentNode, &wg)
	}

	// Wait for all routines to finish
	wg.Wait()

	logrus.WithFields(logrus.Fields{
		"component":    "DESTROY_ENGINE",
		"deploymentId": entDeployment.ID,
	}).Debug("deployment destroyed!")

	// Set the deployment as DESTROYED
	err = entDeployment.Update().
		SetState(deployment.StateDestroyed).
		Exec(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "DESTROY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to update deployment state: %v", err)
		return
	}

	// Query all of the resources
	entResources, err := entDeployment.QueryBlueprint().QueryResources().All(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "DESTROY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to query resources from deployment: %v", err)
		return
	}
	// Query the project
	entProject, err := entDeployment.QueryProject().Only(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "DESTROY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to query project from deployment: %v", err)
		return
	}

	// Subtract out the new usage level
	newUsageCPU := entProject.UsageCPU
	newUsageRAM := entProject.UsageRAM
	newUsageDisk := entProject.UsageDisk
	newUsageNetwork := entProject.UsageNetwork
	newUsageRouter := entProject.UsageRouter
	for _, entResource := range entResources {
		newUsageCPU -= int(entResource.QuotaRequirements.Cpu)
		newUsageRAM -= int(entResource.QuotaRequirements.Ram)
		newUsageDisk -= int(entResource.QuotaRequirements.Disk)
		newUsageNetwork -= int(entResource.QuotaRequirements.Network)
		newUsageRouter -= int(entResource.QuotaRequirements.Router)
	}

	// Update the project quota usage
	err = entProject.Update().
		SetUsageCPU(newUsageCPU).
		SetUsageRAM(newUsageRAM).
		SetUsageDisk(newUsageDisk).
		SetUsageNetwork(newUsageNetwork).
		SetUsageRouter(newUsageRouter).
		Exec(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "DESTROY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to update project quota usage: %v", err)
		return
	}
}

func destroyRoutine(ctx context.Context, cbleServer *providers.CBLEServer, entDeploymentNode *ent.DeploymentNode, wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.WithFields(logrus.Fields{
		"component":        "DESTROY_ENGINE",
		"deploymentNodeId": entDeploymentNode.ID,
	}).Debug("destroy routine starting")

	entResource, err := entDeploymentNode.QueryResource().Only(ctx)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Errorf("failed to query resource from deployment node: %v", err)
		return
	}

	// Auto-mark data nodes as destroyed
	if entResource.Type != resource.TypeResource {
		err = entDeploymentNode.Update().
			SetState(deploymentnode.StateDestroyed).
			Exec(ctx)
		if err != nil {
			// Mark node as failed
			failNode(ctx, entDeploymentNode)
			// Log error
			logrus.WithFields(logrus.Fields{
				"component":        "DESTROY_ENGINE",
				"deploymentNodeId": entDeploymentNode.ID,
			}).Errorf("failed to update node vars and state: %v", err)
		}
		return
	}

	// If the node is not awaiting destruction, return
	if entDeploymentNode.State != deploymentnode.StateToDestroy {
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Debug("node not in state \"to_destroy\"")
		return
	}

	// Set the node's status to CHILD_AWAITING
	err = setStatus(ctx, entDeploymentNode, deploymentnode.StateChildAwaiting)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Error(err)
		return
	}

	logrus.WithFields(logrus.Fields{
		"component":        "DESTROY_ENGINE",
		"deploymentNodeId": entDeploymentNode.ID,
	}).Debug("waiting for children to destroy")

	// Wait for all next nodes to be destroyed
	for {
		// Query all of the undestroyed next nodes (dependencies)
		nextNodes, err := entDeploymentNode.QueryNextNodes().Where(
			deploymentnode.StateNEQ(deploymentnode.StateDestroyed), // Query only undestroyed
		).All(ctx)
		if err != nil {
			// Mark node as failed
			failNode(ctx, entDeploymentNode)
			// Log error
			logrus.WithFields(logrus.Fields{
				"component":        "DESTROY_ENGINE",
				"deploymentNodeId": entDeploymentNode.ID,
			}).Errorf("failed to query uncompleted next nodes from node: %v", err)
			return
		}

		// If no undestroyed next nodes, move on to destroying this node
		if len(nextNodes) == 0 {
			break
		}

		// Wait 1 second in-between checking next node statuses
		time.Sleep(time.Second)
	}

	logrus.WithFields(logrus.Fields{
		"component":        "DESTROY_ENGINE",
		"deploymentNodeId": entDeploymentNode.ID,
	}).Debug("children completed")

	// Set the node's status to IN_PROGRESS
	err = setStatus(ctx, entDeploymentNode, deploymentnode.StateInProgress)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Error(err)
		return
	}

	// Query the provider from the resource
	entProvider, err := entDeploymentNode.QueryDeployment().
		QueryBlueprint().
		QueryProvider().
		Only(ctx)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Errorf("failed to query provider from node: %v", err)
		return
	}

	logrus.WithFields(logrus.Fields{
		"component":        "DESTROY_ENGINE",
		"deploymentNodeId": entDeploymentNode.ID,
	}).Debug("destroying resource...")

	// Have the provider destroy the resource
	reply, err := cbleServer.DestroyResource(ctx, entProvider, entDeploymentNode)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Errorf("failed to destroy resource: %v", err)
		return
	}
	if !reply.Success {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Errorf("failed to destroy resource: %s", *reply.Error)
		return
	}

	logrus.WithFields(logrus.Fields{
		"component":        "DESTROY_ENGINE",
		"deploymentNodeId": entDeploymentNode.ID,
	}).Debug("destroyed resource successfully!")

	// Update the vars and state on success
	err = entDeploymentNode.Update().
		SetVars(reply.UpdatedVars).
		SetState(deploymentnode.StateDestroyed).
		Exec(ctx)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithFields(logrus.Fields{
			"component":        "DESTROY_ENGINE",
			"deploymentNodeId": entDeploymentNode.ID,
		}).Errorf("failed to update node vars and state: %v", err)
		return
	}
}
