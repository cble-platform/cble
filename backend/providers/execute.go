package providers

import (
	"context"
	"fmt"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sirupsen/logrus"
)

// Blocking call which executes the given provider and handles shutdown when requested
func ExecuteProvider(ctx context.Context, cacheDir string, entProvider *ent.Provider, shutdown chan bool) {
	providerRepoPath := path.Join(cacheDir, entProvider.ID.String())
	providerMetadata, err := ParseMetadata(providerRepoPath)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Errorf("failed to parse provider metadata: %v", err)
		return
	}

	switch providerMetadata.Type {
	case TypeDocker:
		execDockerProvider(ctx, providerMetadata, entProvider, shutdown)
	case TypeShell:
		execShellProvider(ctx, providerRepoPath, providerMetadata, entProvider, shutdown)
	default:
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Errorf("unknown provider type %s", providerMetadata.Type)
	}
}

func execDockerProvider(ctx context.Context, metadata *ProviderMetadata, entProvider *ent.Provider, shutdown chan bool) {
	if metadata.DockerMeta == nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Errorf("\"docker\" block not found in provider metadata, please provide one")
		return
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer dockerClient.Close()

	// Generate the provider command (with ID as last positional param)
	providerCmd := strings.Split(metadata.DockerMeta.Command, " ")
	providerCmd = append(providerCmd, entProvider.ID.String())

	// Create the docker container
	resp, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image: generateDockerImageName(metadata),
		Cmd:   providerCmd,
	}, &container.HostConfig{
		NetworkMode: container.NetworkMode("host"), // Use host networking
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/tmp",
				Target: "/tmp",
			},
		},
	}, &network.NetworkingConfig{}, &v1.Platform{}, fmt.Sprintf("cble-provider-%s", entProvider.ID.String()))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Errorf("failed to create container: %v", err)
		return
	}

	// Start the container in the background
	if err := dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Errorf("failed to start container: %v", err)
		return
	}

	// Defer the container shutdown
	defer func() {
		// Use background context since upper context is cancelled
		ctx := context.Background()

		if err := dockerClient.ContainerStop(ctx, resp.ID, container.StopOptions{}); err != nil {
			logrus.WithFields(logrus.Fields{
				"component":  "PROVIDER_ENGINE",
				"providerId": entProvider.ID,
			}).Errorf("failed to stop container: %v", err)
			return
		}
		if err := dockerClient.ContainerRemove(ctx, resp.ID, container.RemoveOptions{}); err != nil {
			logrus.WithFields(logrus.Fields{
				"component":  "PROVIDER_ENGINE",
				"providerId": entProvider.ID,
			}).Errorf("failed to remove container: %v", err)
			return
		}
	}()

	for {
		select {
		case <-ctx.Done():
			logrus.WithFields(logrus.Fields{
				"component":  "PROVIDER_ENGINE",
				"providerId": entProvider.ID,
			}).Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			return
		case <-shutdown:
			logrus.WithFields(logrus.Fields{
				"component":  "PROVIDER_ENGINE",
				"providerId": entProvider.ID,
			}).Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			return
		}
	}
}

func execShellProvider(ctx context.Context, providerRepoPath string, metadata *ProviderMetadata, entProvider *ent.Provider, shutdown chan bool) {
	if metadata.ShellMeta == nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Errorf("\"shell\" block not found in provider metadata, please provide one with a build_cmd")
		return
	}

	// Start the provider with the provider ID as argument
	cmd := exec.Command(metadata.ShellMeta.ExecCommand, entProvider.ID.String())
	cmd.Dir = providerRepoPath
	if err := cmd.Start(); err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Errorf("failed to run provider server: failed to start provider: %v", err)
		return
	}

	// Defer the process shutdown
	defer func() {
		cmd.Process.Signal(syscall.SIGTERM)
	}()

	for {
		select {
		case <-ctx.Done():
			logrus.WithFields(logrus.Fields{
				"component":  "PROVIDER_ENGINE",
				"providerId": entProvider.ID,
			}).Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			return
		case <-shutdown:
			logrus.WithFields(logrus.Fields{
				"component":  "PROVIDER_ENGINE",
				"providerId": entProvider.ID,
			}).Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			return
		}
	}
}
