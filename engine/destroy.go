package engine

import (
	"context"
	"sync"
	"time"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
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
		logrus.Errorf("failed to update deployment state: %v", err)
		return
	}

	// Set all of the nodes to TO_DESTROY state (which aren't destroyed)
	err = client.DeploymentNode.Update().
		Where(
			deploymentnode.And(
				deploymentnode.HasDeploymentWith(deployment.IDEQ(entDeployment.ID)),
				deploymentnode.StateNEQ(deploymentnode.StateDestroyed),
			),
		).
		SetState(deploymentnode.StateToDestroy).
		Exec(ctx)
	if err != nil {
		logrus.Errorf("failed to set deployment nodes to TO_DESTROY: %v", err)
	}

	// Query all of the deployment nodes
	entDeploymentNodes, err := entDeployment.QueryDeploymentNodes().All(ctx)
	if err != nil {
		logrus.Errorf("failed to query deployment nodes: %v", err)
		return
	}

	var wg sync.WaitGroup

	// Spawn destroyRoutine's for all root nodes
	for _, entDeploymentNode := range entDeploymentNodes {
		wg.Add(1)
		go destroyRoutine(ctx, client, cbleServer, entDeploymentNode, &wg)
	}

	// Wait for all routines to finish
	wg.Wait()

	logrus.Debug("deployment destroyed!")

	// Set the deployment as DESTROYED
	err = entDeployment.Update().
		SetState(deployment.StateDestroyed).
		Exec(ctx)
	if err != nil {
		logrus.Errorf("failed to update deployment state: %v", err)
		return
	}
}

func destroyRoutine(ctx context.Context, client *ent.Client, cbleServer *providers.CBLEServer, entDeploymentNode *ent.DeploymentNode, wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.WithField("node", entDeploymentNode.ID).Debug("destroy routine starting")

	// If the node is not awaiting destruction, return
	if entDeploymentNode.State != deploymentnode.StateToDestroy {
		logrus.WithField("node", entDeploymentNode.ID).Debug("node not in state \"to_destroy\"")
		return
	}

	// Set the node's status to CHILD_AWAITING
	err := setStatus(ctx, entDeploymentNode, deploymentnode.StateChildAwaiting)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithField("node", entDeploymentNode.ID).Error(err)
		return
	}

	logrus.WithField("node", entDeploymentNode.ID).Debug("waiting for children to destroy")

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
			logrus.WithField("node", entDeploymentNode.ID).Errorf("failed to query uncompleted next nodes from node: %v", err)
			return
		}

		// If no undestroyed next nodes, move on to destroying this node
		if len(nextNodes) == 0 {
			break
		}

		// Wait 1 second in-between checking next node statuses
		time.Sleep(time.Second)
	}

	logrus.WithField("node", entDeploymentNode.ID).Debug("children completed")

	// Set the node's status to IN_PROGRESS
	err = setStatus(ctx, entDeploymentNode, deploymentnode.StateInProgress)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.WithField("node", entDeploymentNode.ID).Error(err)
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
		logrus.WithField("node", entDeploymentNode.ID).Errorf("failed to query provider from node: %v", err)
		return
	}

	logrus.WithField("node", entDeploymentNode.ID).Debug("destroying resource...")

	// Have the provider destroy the resource
	reply, err := cbleServer.DestroyResource(ctx, entProvider, entDeploymentNode)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.Errorf("failed to destroy resource: %v", err)
		return
	}
	if !reply.Success {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.Errorf("failed to destroy resource: %s", *reply.Errors)
		return
	}

	logrus.WithField("node", entDeploymentNode.ID).Debug("destroyed resource successfully!")

	// Update the vars and state on success
	err = entDeploymentNode.Update().
		SetVars(reply.UpdatedVars).
		SetState(deploymentnode.StateDestroyed).
		Exec(ctx)
	if err != nil {
		// Mark node as failed
		failNode(ctx, entDeploymentNode)
		// Log error
		logrus.Errorf("failed to update node vars and state: %v", err)
		return
	}
}
