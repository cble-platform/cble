package engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/sirupsen/logrus"
)

// Run in a go routine, runs the deployment with a provider
func StartDeployment(client *ent.Client, cbleServer *providers.CBLEServer, entDeployment *ent.Deployment) {
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

	// Query all of the deployment nodes
	entDeploymentNodes, err := entDeployment.QueryDeploymentNodes().All(ctx)
	if err != nil {
		logrus.Errorf("failed to query deployment nodes: %v", err)
		return
	}

	var wg sync.WaitGroup

	// Spawn deployRoutine's for all root nodes
	for _, deploymentNode := range entDeploymentNodes {
		go deployRoutine(ctx, client, cbleServer, deploymentNode, &wg)
	}
}

func deployRoutine(ctx context.Context, client *ent.Client, cbleServer *providers.CBLEServer, deploymentNode *ent.DeploymentNode, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	logrus.WithField("node", deploymentNode.ID).Debug("deploy routine starting")

	// If the node is not awaiting deployment, return
	if deploymentNode.State != deploymentnode.StateAwaiting {
		logrus.WithField("node", deploymentNode.ID).Debug("node not in state \"awaiting\"")
		return
	}

	// Set the node's status to PARENT_AWAITING
	err := setStatus(ctx, deploymentNode, deploymentnode.StateParentAwaiting)
	if err != nil {
		// Mark node as failed
		failNode(ctx, deploymentNode)
		// Log error
		logrus.WithField("node", deploymentNode.ID).Error(err)
		return
	}

	logrus.WithField("node", deploymentNode.ID).Debug("waiting for parents to complete")

	// Wait for all prev nodes to be completed
	for {
		// Query all of the uncompleted previous nodes (dependencies)
		prevNodes, err := deploymentNode.QueryPrevNodes().Where(
			deploymentnode.StateNEQ(deploymentnode.StateComplete), // Query only uncompleted
		).All(ctx)
		if err != nil {
			// Mark node as failed
			failNode(ctx, deploymentNode)
			// Log error
			logrus.WithField("node", deploymentNode.ID).Errorf("failed to query uncompleted prev nodes from node: %v", err)
			return
		}

		// If no uncompleted prev nodes, move on to deploying this node
		if len(prevNodes) == 0 {
			break
		}

		for _, prevNode := range prevNodes {
			// If the prev node is failed, mark this node as tainted and stop
			if prevNode.State == deploymentnode.StateFailed {
				err = setStatus(ctx, deploymentNode, deploymentnode.StateTainted)
				if err != nil {
					logrus.WithField("node", deploymentNode.ID).Error(err)
				}
				return
			}
		}

		// Wait 1 second in-between checking previous node statuses
		time.Sleep(time.Second)
	}

	logrus.WithField("node", deploymentNode.ID).Debug("parents completed")

	// Set the node's status to IN_PROGRESS
	err = setStatus(ctx, deploymentNode, deploymentnode.StateInProgress)
	if err != nil {
		// Mark node as failed
		failNode(ctx, deploymentNode)
		// Log error
		logrus.WithField("node", deploymentNode.ID).Error(err)
		return
	}

	// Query the provider from the resource
	entProvider, err := deploymentNode.QueryDeployment().
		QueryBlueprint().
		QueryProvider().
		Only(ctx)
	if err != nil {
		// Mark node as failed
		failNode(ctx, deploymentNode)
		// Log error
		logrus.WithField("node", deploymentNode.ID).Errorf("failed to query provider from node: %v", err)
		return
	}

	logrus.WithField("node", deploymentNode.ID).Debug("deploying resource...")

	// Have the provider deploy the resource
	reply, err := cbleServer.DeployResource(ctx, entProvider, deploymentNode)
	if err != nil {
		// Mark node as failed
		failNode(ctx, deploymentNode)
		// Log error
		logrus.Errorf("failed to deploy resource: %v", err)
		return
	}
	if !reply.Success {
		// Mark node as failed
		failNode(ctx, deploymentNode)
		// Log error
		logrus.Errorf("failed to deploy resource: %s", *reply.Errors)
		return
	}

	logrus.WithField("node", deploymentNode.ID).Debug("deployed resource successfully!")

	// Update the vars and state on success
	err = deploymentNode.Update().
		SetVars(reply.UpdatedVars).
		SetState(deploymentnode.StateComplete).
		Exec(ctx)
	if err != nil {
		// Mark node as failed
		failNode(ctx, deploymentNode)
		// Log error
		logrus.Errorf("failed to update node vars and state: %v", err)
		return
	}
}

func failNode(ctx context.Context, entDeploymentNode *ent.DeploymentNode) {
	// Mark node as failed
	err := setStatus(ctx, entDeploymentNode, deploymentnode.StateFailed)
	if err != nil {
		logrus.WithField("node", entDeploymentNode.ID).Error(err)
	}
}

func setStatus(ctx context.Context, entDeploymentNode *ent.DeploymentNode, state deploymentnode.State) error {
	err := entDeploymentNode.Update().
		SetState(state).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update node state to \"%s\": %v", state, err)
	}
	return nil
}
