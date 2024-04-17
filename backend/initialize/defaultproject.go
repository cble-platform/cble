package initialize

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/project"
)

func InitDefaultProject(ctx context.Context, client *ent.Client, cbleConfig *config.Config) (*ent.Project, error) {
	// Ensure the default project exists
	cbleDefaultProject, err := client.Project.Query().Where(
		project.NameEQ(cbleConfig.Initialization.DefaultProject),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("failed to query for default project: %v", err)
	} else if cbleDefaultProject == nil {
		// If it doesn't exist, make it
		cbleDefaultProject, err = client.Project.Create().
			SetName(cbleConfig.Initialization.DefaultProject).
			SetQuotaCPU(cbleConfig.ProjectDefaults.QuotaCPU).
			SetQuotaRAM(cbleConfig.ProjectDefaults.QuotaRAM).
			SetQuotaDisk(cbleConfig.ProjectDefaults.QuotaDisk).
			SetQuotaNetwork(cbleConfig.ProjectDefaults.QuotaNetwork).
			SetQuotaRouter(cbleConfig.ProjectDefaults.QuotaRouter).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create default project: %v", err)
		}
	}

	return cbleDefaultProject, nil
}
