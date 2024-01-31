package blueprintengine

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble-backend/blueprintengine/models"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/blueprint"
	"github.com/cble-platform/cble-backend/ent/resource"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func LoadResources(ctx context.Context, client *ent.Client, entBlueprint *ent.Blueprint) error {
	// Parse the template into our generic objects
	var parsedBlueprint *models.Blueprint
	if err := yaml.Unmarshal(entBlueprint.BlueprintTemplate, &parsedBlueprint); err != nil {
		return fmt.Errorf("failed to parse blueprint: %v", err)
	}

	entResources := make([]*ent.Resource, 0)
	referencedKeys := make([]string, 0)

	// Ensure all the objects have resources created
	for key, object := range parsedBlueprint.Objects {
		logrus.Debugf("Loading resource %s", key)

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
			entResource, err = entResource.Update().
				SetObject(object).
				Save(ctx)
			if err != nil {
				return err
			}
		} else {
			// Otherwise, create it
			entResource, err = client.Resource.Create().
				SetKey(key).
				SetObject(object).
				SetBlueprint(entBlueprint).
				Save(ctx)
			if err != nil {
				return err
			}
		}

		// Add it to the array of resources
		entResources = append(entResources, entResource)
		referencedKeys = append(referencedKeys, key)
	}

	// Update all the resource dependencies
	for _, entResource := range entResources {
		logrus.Debugf("Updating resource dependency for resource %s", entResource.ID)

		// Query all of the dependencies of this resource
		entResourceDependencies, err := client.Resource.Query().Where(
			resource.KeyIn(entResource.Object.DependsOn...),
		).All(ctx)
		if err != nil {
			return err
		}

		// Update the resource with dependencies
		entResource, err = entResource.Update().
			ClearDependsOn().
			Save(ctx)
		if err != nil {
			return err
		}
		err = entResource.Update().
			AddDependsOn(entResourceDependencies...).
			Exec(ctx)
		if err != nil {
			return err
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
	logrus.Debugf("Deleted %d dangling resources", deleteCount)

	return nil
}
