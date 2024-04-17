package engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/cble-platform/cble/backend/ent"
	"github.com/cble-platform/cble/backend/ent/deployment"
	"github.com/cble-platform/cble/backend/ent/deploymentnode"
	"github.com/cble-platform/cble/backend/ent/resource"
	"github.com/cble-platform/cble/backend/providers"
	"github.com/sirupsen/logrus"
)

// AutoSuspendDeploymentWatchdog is designed to be executed as a go routine which searches for all deployments powered on which haven't been accessed in a time period
func AutoSuspendDeploymentWatchdog(ctx context.Context, client *ent.Client, cbleServer *providers.CBLEServer, suspendTime time.Duration) {
	// Query for stale deployments every 30 minutes
	ticker := time.NewTicker(30 * time.Minute)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			logrus.WithField("component", "AUTO_SUSPEND").Debug("auto-suspending deployments")

			// Find all deployments which are considered "stale"
			staleDeployments, err := client.Deployment.Query().Where(
				deployment.StateEQ(deployment.StateComplete),                         // Is currently in the COMPLETE (powered on) state
				deployment.LastAccessedLTE(time.Now().Add(-suspendTime*time.Minute)), // last accessed at least suspendTime ago
			).All(ctx)
			if err != nil {
				logrus.WithField("component", "AUTO_SUSPEND").Errorf("failed to query stale deployments for auto-suspend: %v", err)
				continue
			}

			var wg sync.WaitGroup

			// Power them all off in parallel
			for _, staleDeployment := range staleDeployments {
				wg.Add(1)
				go powerOffDeployment(ctx, cbleServer, staleDeployment, &wg)
			}

			wg.Wait()
		}
	}
}

func powerOffDeployment(ctx context.Context, cbleServer *providers.CBLEServer, entDeployment *ent.Deployment, wg *sync.WaitGroup) {
	defer wg.Done()

	// Get the provider
	entProvider, err := entDeployment.QueryBlueprint().QueryProvider().Only(ctx)
	if err != nil {
		return
	}
	// Get all of the deployment nodes which support power feature
	entDeploymentNodes, err := entDeployment.QueryDeploymentNodes().Where(
		deploymentnode.HasResourceWith(func(s *sql.Selector) {
			s.Where(sqljson.ValueEQ(resource.FieldFeatures, true, sqljson.Path("power"))) // Find where has { "power": true, ... }
		}),
	).All(ctx)
	if err != nil {
		logrus.WithField("component", "AUTO_SUSPEND").Errorf("failed to query nodes for deployment %s: %v", entDeployment.ID, err)
		return
	}

	var nodeWg sync.WaitGroup
	var reply *provider.ResourcePowerReply

	for _, entDeploymentNode := range entDeploymentNodes {
		nodeWg.Add(1)
		go func(nodeWg *sync.WaitGroup, entDeploymentNode *ent.DeploymentNode) {
			defer nodeWg.Done()
			// Update the resource power state
			reply, err = cbleServer.ResourcePower(ctx, entProvider, entDeploymentNode, provider.PowerState_OFF)
			if reply != nil && !reply.Success {
				if reply.Error != nil {
					err = fmt.Errorf("failed to power off node: %v", reply.Error)
				} else {
					err = fmt.Errorf("failed to power off node: unknown error")
				}
			}
		}(&nodeWg, entDeploymentNode)
	}

	// Wait for all of the resources to power down
	nodeWg.Wait()

	if err != nil {
		logrus.WithField("component", "AUTO_SUSPEND").Errorf("encountered at least one error while auto-suspending deployment %s: %v", entDeployment.ID, err)
	} else {
		// If we successfully applied all power states, update the deployment state
		err = entDeployment.Update().SetState(deployment.StateSuspended).Exec(ctx)
		if err != nil {
			logrus.WithField("component", "AUTO_SUSPEND").Errorf("failed to update state for deployment %s: %v", entDeployment.ID, err)
			return
		}
	}
}

func AutoDestroyExpiredDeploymentWatchdog(ctx context.Context, client *ent.Client, cbleServer *providers.CBLEServer) {
	// Query for expired deployments every 30 minutes
	ticker := time.NewTicker(30 * time.Minute)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			logrus.WithField("component", "AUTO_DESTROY").Debug("destroying expired deployments")

			// Find all deployments which are expired
			expiredDeployments, err := client.Deployment.Query().Where(
				deployment.StateNEQ(deployment.StateDestroyed),  // Is not already destroyed
				deployment.StateNEQ(deployment.StateInProgress), // Is not currently destroying
				deployment.ExpiresAtLTE(time.Now()),             // Expiry time has passed
			).All(ctx)
			if err != nil {
				logrus.WithField("component", "AUTO_DESTROY").Errorf("failed to query expired deployments for auto-destructions: %v", err)
				continue
			}

			// Destroy the expired deployments
			for _, expiredDeployment := range expiredDeployments {
				go StartDestroy(client, cbleServer, expiredDeployment)
			}
		}
	}
}
