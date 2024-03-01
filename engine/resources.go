package engine

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble-backend/engine/models"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/blueprint"
	"github.com/cble-platform/cble-backend/ent/resource"
	"github.com/cble-platform/cble-backend/providers"
	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func LoadResources(ctx context.Context, client *ent.Client, cbleServer *providers.CBLEServer, entBlueprint *ent.Blueprint) error {
	// Parse the template into our generic objects
	var parsedBlueprint *models.Blueprint
	if err := yaml.Unmarshal(entBlueprint.BlueprintTemplate, &parsedBlueprint); err != nil {
		return fmt.Errorf("failed to parse blueprint: %v", err)
	}

	entResources := make([]*ent.Resource, 0)
	referencedKeys := make([]string, 0)
	resourceMap := make(map[string]*ent.Resource, 0)

	// Ensure all the objects have resources created
	for key, object := range parsedBlueprint.Objects {
		if object.Resource == nil && object.Data == nil {
			return fmt.Errorf("object %s has no resource or data type", key)
		}

		logrus.WithFields(logrus.Fields{
			"component":   "BLUEPRINT_ENGINE",
			"blueprintId": entBlueprint.ID,
		}).Debugf("Loading resource %s", key)

		// Check if the resource already exists
		entResource, err := client.Resource.Query().Where(
			resource.And(
				resource.HasBlueprintWith(blueprint.IDEQ(entBlueprint.ID)),
				resource.KeyEQ(key),
			),
		).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		// If the resource exists, update it
		if entResource != nil {
			entResourceUpdate := entResource.Update().
				SetObject(object).
				ClearDependsOn() // Clear the dependencies for updating later
			// Set the resource type from "resource" or "data"
			if object.Resource != nil {
				entResourceUpdate = entResourceUpdate.SetType(resource.TypeResource).SetResourceType(*object.Resource)
			} else if object.Data != nil {
				entResourceUpdate = entResourceUpdate.SetType(resource.TypeData).SetResourceType(*object.Data)
			}
			entResource, err = entResourceUpdate.Save(ctx)
			if err != nil {
				return err
			}
		} else {
			// Otherwise, create it
			entResourceCreate := client.Resource.Create().
				SetKey(key).
				SetObject(object).
				SetBlueprint(entBlueprint)
			if object.Resource != nil {
				entResourceCreate = entResourceCreate.SetType(resource.TypeResource).SetResourceType(*object.Resource)
			} else if object.Data != nil {
				entResourceCreate = entResourceCreate.SetType(resource.TypeData).SetResourceType(*object.Data)
			}
			entResource, err = entResourceCreate.Save(ctx)
			if err != nil {
				return err
			}
		}

		// Add it to the array of resources
		entResources = append(entResources, entResource)
		referencedKeys = append(referencedKeys, key)
		resourceMap[key] = entResource
	}

	// Request the resources metadata from the provider
	entProvider, err := entBlueprint.QueryProvider().Only(ctx)
	if err != nil {
		return fmt.Errorf("failed to query provider from blueprint: %v", err)
	}
	metadataReply, err := cbleServer.ExtractResourceMetadata(ctx, entProvider, entResources)
	if err != nil {
		return fmt.Errorf("failed to call extract resource metadata: %v", err)
	}
	if !metadataReply.Success {
		if metadataReply.Error != nil {
			return fmt.Errorf("failed extract resource metadata: %s", *metadataReply.Error)
		}
		return fmt.Errorf("failed to extract resource metadata: unknown error")
	}

	// Store all metadata of resources in the database
	for resourceKey, resourceMetadata := range metadataReply.Metadata {
		entResource, ok := resourceMap[resourceKey]
		if !ok {
			return fmt.Errorf("failed to pull resource %s from resource map: %v", resourceKey, err)
		}

		// Update resource features
		if resourceMetadata.Features != nil {
			err = entResource.Update().
				SetFeatures(pgrpc.Features{ // Need to create new struct b/c otherwise copies lock value
					Power:   resourceMetadata.Features.Power,
					Console: resourceMetadata.Features.Console,
				}).
				SetQuotaRequirements(pgrpc.QuotaRequirements{
					Cpu:     resourceMetadata.QuotaRequirements.Cpu,
					Ram:     resourceMetadata.QuotaRequirements.Ram,
					Disk:    resourceMetadata.QuotaRequirements.Disk,
					Router:  resourceMetadata.QuotaRequirements.Router,
					Network: resourceMetadata.QuotaRequirements.Network,
				}).
				Exec(ctx)
			if err != nil {
				return fmt.Errorf("failed to update features for resource %s: %v", resourceKey, err)
			}
		}

		// Update all the resource dependencies
		for _, dependencyResourceKey := range resourceMetadata.DependsOnKeys {
			entResourceDependency, ok := resourceMap[dependencyResourceKey]
			if !ok {
				return fmt.Errorf("failed to pull resource dependency %s from resource map: %v", dependencyResourceKey, err)
			}
			logrus.WithFields(logrus.Fields{
				"component":   "BLUEPRINT_ENGINE",
				"blueprintId": entBlueprint.ID,
			}).Debugf("Updating resource dependency for resource %s", entResource.ID)

			// Update the resource with dependencies
			err = entResource.Update().
				AddDependsOn(entResourceDependency).
				Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	// Delete all resources which no longer exist
	deleteCount, err := client.Resource.Delete().Where(
		resource.And(
			resource.HasBlueprintWith(blueprint.IDEQ(entBlueprint.ID)), // Resources attached to this blueprint
			resource.KeyNotIn(referencedKeys...),                       // Which are no longer referenced
		),
	).Exec(ctx)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"component":   "BLUEPRINT_ENGINE",
		"blueprintId": entBlueprint.ID,
	}).Debugf("Deleted %d dangling resources", deleteCount)

	return nil
}
