package engine

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/sirupsen/logrus"
)

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
