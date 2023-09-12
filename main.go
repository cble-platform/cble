package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cble-platform/backend/config"
	"github.com/cble-platform/backend/ent"
	"github.com/cble-platform/backend/ent/group"
	"github.com/cble-platform/backend/ent/user"
	"github.com/cble-platform/backend/internal"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"golang.org/x/crypto/bcrypt"
)

var cfgFile string

func main() {
	// CLI flags
	pflag.StringVar(&cfgFile, "config", "", "the path to the config file")
	pflag.Parse()

	fmt.Println(internal.Logo())
	cbleConfig, err := config.LoadConfig(cfgFile)
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()

	//-----//
	// ENT //
	//-----//

	pgPort := 5432
	if cbleConfig.Database.Port != nil {
		pgPort = *cbleConfig.Database.Port
	}
	pgDatabase := "cble"
	if cbleConfig.Database.Database != nil {
		pgDatabase = *cbleConfig.Database.Database
	}
	pgSslMode := "disable"
	if cbleConfig.Database.SSL != nil && *cbleConfig.Database.SSL {
		pgSslMode = "require"
	}
	pgConnStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cbleConfig.Database.Host,
		pgPort,
		cbleConfig.Database.Username,
		pgDatabase,
		cbleConfig.Database.Password,
		pgSslMode,
	)
	client, err := ent.Open("postgres", pgConnStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//---------------//
	// Default Admin //
	//---------------//

	// Ensure the built-in admin group exists
	cbleAdminGroup, err := client.Group.Query().Where(
		group.NameEQ(cbleConfig.Initialization.AdminGroup),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		logrus.Fatalf("failed to query for default admin group: %v", err)
	} else if cbleAdminGroup == nil {
		// If it doesn't exist, make it
		cbleAdminGroup, err = client.Group.Create().
			SetName(cbleConfig.Initialization.AdminGroup).
			Save(ctx)
		if err != nil {
			logrus.Fatalf("failed to create default admin group: %v", err)
		}
	}

	// Check if there are any admins in existence with the built in admin group
	if anyAdminExists, err := client.User.Query().Where(
		user.And(
			user.HasGroupsWith(
				group.NameEQ(cbleConfig.Initialization.AdminGroup),
			),
		),
	).Exist(ctx); err != nil {
		logrus.Fatalf("failed to query for existing admins: %v", err)
	} else if !anyAdminExists {
		// If there are no admins in admin group, check if the default one exists
		defaultAdminExists, err := client.User.Query().Where(
			user.And(
				user.UsernameEQ(cbleConfig.Initialization.DefaultAdmin.Username),
				user.HasGroupsWith(
					group.NameEQ(cbleConfig.Initialization.AdminGroup),
				),
			),
		).Exist(ctx)
		if err != nil {
			logrus.Fatalf("failed to query for default admin: %v", err)
		}
		if !defaultAdminExists {
			// If the default one doesn't exist, make it

			// Hash the default password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cbleConfig.Initialization.DefaultAdmin.Password), 8)
			if err != nil {
				logrus.Fatalf("failed to hash default admin password: %v", err)
			}
			// Create the default admin
			_, err = client.User.Create().
				SetUsername(cbleConfig.Initialization.DefaultAdmin.Username).
				SetPassword(string(hashedPassword)).
				SetEmail(cbleConfig.Initialization.DefaultAdmin.Email).
				SetFirstName(cbleConfig.Initialization.DefaultAdmin.FirstName).
				SetLastName(cbleConfig.Initialization.DefaultAdmin.LastName).
				AddGroups(cbleAdminGroup).
				Save(ctx)
			if err != nil {
				logrus.Fatalf("failed to create default admin: %v", err)
			}
			logrus.Info("Created default admin user")
		}
	}
}
