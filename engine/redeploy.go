package engine

import (
	"context"
	"fmt"
	"sync"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Run in a go routine, runs the deployment with a provider
func StartRedeploy(client *ent.Client, cbleServer *providers.CBLEServer, entDeployment *ent.Deployment, nodeIdsToRedeploy []uuid.UUID) {
	// Create new context for deployment
	ctx := context.Background()

	// Query all of the deployment nodes to redeploy
	entDeploymentNodes, err := entDeployment.QueryDeploymentNodes().Where(
		deploymentnode.IDIn(nodeIdsToRedeploy...),
	).All(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "REDEPLOY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to query deployment nodes: %v", err)
		return
	}

	// Set the deployment as IN_PROGRESS
	entDeployment, err = entDeployment.Update().
		SetState(deployment.StateInProgress).
		Save(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "REDEPLOY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to update deployment state: %v", err)
		return
	}

	// Update all of the nodes and children to to_destroy
	for _, deploymentNode := range entDeploymentNodes {
		err = markRedeployChildren(ctx, client, deploymentNode, deploymentnode.StateToDestroy)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"component":    "REDEPLOY_ENGINE",
				"deploymentId": entDeployment.ID,
			}).Errorf("failed to mark nodes for redeploy: %v", err)
			return
		}
	}

	// Query all of the deployment nodes marked for redeploy
	entDeploymentNodes, err = entDeployment.QueryDeploymentNodes().All(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "REDEPLOY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to query deployment nodes: %v", err)
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

	logrus.WithFields(logrus.Fields{
		"component":    "REDEPLOY_ENGINE",
		"deploymentId": entDeployment.ID,
	}).Debug("deployment destroyed!")

	// Set all destroyed resources as TO_DEPLOY
	err = client.DeploymentNode.Update().
		SetState(deploymentnode.StateToDeploy).
		Where(
			deploymentnode.And(
				deploymentnode.HasDeploymentWith(deployment.IDEQ(entDeployment.ID)),
				deploymentnode.StateEQ(deploymentnode.StateDestroyed),
			),
		).
		Exec(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "REDEPLOY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to mark all destroyed nodes for redeployment: %v", err)
		return
	}

	// Query all of the deployment nodes marked for redeploy
	entDeploymentNodes, err = entDeployment.QueryDeploymentNodes().All(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "REDEPLOY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to query deployment nodes: %v", err)
		return
	}

	// Spawn deployRoutine's for all root nodes
	for _, deploymentNode := range entDeploymentNodes {
		wg.Add(1)
		go deployRoutine(ctx, client, cbleServer, deploymentNode, &wg)
	}

	// Wait for all routines to finish
	wg.Wait()

	logrus.WithFields(logrus.Fields{
		"component":    "REDEPLOY_ENGINE",
		"deploymentId": entDeployment.ID,
	}).Debug("deployment successful!")

	// Set the deployment as COMPLETE
	err = entDeployment.Update().
		SetState(deployment.StateComplete).
		Exec(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":    "REDEPLOY_ENGINE",
			"deploymentId": entDeployment.ID,
		}).Errorf("failed to update deployment state: %v", err)
		return
	}
}

func markRedeployChildren(ctx context.Context, client *ent.Client, deploymentNode *ent.DeploymentNode, state deploymentnode.State) error {
	// Ignore this node if we've already visited it
	if deploymentNode.State == state {
		return nil
	}

	// Set the node's status to new state
	err := setStatus(ctx, deploymentNode, state)
	if err != nil {
		// Log error
		return fmt.Errorf("failed to set node status to to_redeploy: %v", err)
	}

	nextNodes, err := deploymentNode.QueryNextNodes().All(ctx)
	if err != nil {
		return err
	}

	// Mark all of this node's children as new state as well
	for _, nextNode := range nextNodes {
		err = markRedeployChildren(ctx, client, nextNode, state)
		if err != nil {
			return err
		}
	}

	return nil
}
