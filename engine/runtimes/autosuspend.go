package runtimes

import (
	"context"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/ent/resource"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/sirupsen/logrus"
)

// DeploymentAutoSuspendWatchdog is designed to be executed as a go routine which searches for all deployments powered on which haven't been accessed in a time period
func DeploymentAutoSuspendWatchdog(ctx context.Context, client *ent.Client, cbleServer *providers.CBLEServer, suspendTime uint) {
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
				deployment.StateEQ(deployment.StateComplete),                                            // Is currently in the COMPLETE (powered on) state
				deployment.LastAccessedLTE(time.Now().Add(-(time.Duration(suspendTime) * time.Minute))), // last accessed at least suspendTime ago
			).All(ctx)
			if err != nil {
				logrus.WithField("component", "AUTO_SUSPEND").Errorf("failed to query stale deployments for auto-suspend: %v", err)
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
