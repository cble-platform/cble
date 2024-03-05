package providers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/sirupsen/logrus"
)

func BuildProvider(providerRepoPath string, entProvider *ent.Provider) error {
	providerMetadata, err := ParseMetadata(providerRepoPath)
	if err != nil {
		return fmt.Errorf("failed to parse provider metadata: %v", err)
	}

	switch providerMetadata.Type {
	case TypeDocker:
		return buildDockerProvider(providerRepoPath, providerMetadata)
	case TypeShell:
		return buildShellProvider(providerRepoPath, providerMetadata)
	default:
		return fmt.Errorf("unknown provider type %s", providerMetadata.Type)
	}
}

func buildDockerProvider(providerRepoPath string, metadata *ProviderMetadata) error {
	if metadata.DockerMeta == nil {
		return fmt.Errorf("\"docker\" block not found in provider metadata, please provide one")
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer dockerClient.Close()

	// Generate docker build context
	providerSrcTar, err := archive.TarWithOptions(providerRepoPath, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}

	logrus.Debug(providerRepoPath)

	// Build the image
	resp, err := dockerClient.ImageBuild(context.Background(), providerSrcTar, types.ImageBuildOptions{
		Dockerfile: metadata.DockerMeta.Dockerfile,
		Remove:     true,
		Tags:       []string{generateDockerImageName(metadata)},
	})
	if err != nil {
		return fmt.Errorf("failed to execute docker build: %v", err)
	}
	defer resp.Body.Close()

	// Get the last line of the output
	var lastLine string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	// Check if last line was an error message
	errLine := &DockerErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return fmt.Errorf("failed build docker image: %v", errLine.Error)
	}

	// Check if we encountered any scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan docker build output: %v", err)
	}

	logrus.WithField("component", "PROVIDER_BUILD").Debugf("Build docker image for %s", providerRepoPath)
	return nil
}

func buildShellProvider(providerRepoPath string, metadata *ProviderMetadata) error {
	if metadata.ShellMeta == nil {
		return fmt.Errorf("\"shell\" block not found in provider metadata, please provide one with a build_cmd")
	}

	// Build the provider using the provided command
	cmd := exec.Command("sh", "-c", metadata.ShellMeta.BuildCommand)
	cmd.Dir = providerRepoPath
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("command exited with status %d. output: %s", exiterr.ExitCode(), cmdOutput)
		}
		return fmt.Errorf("command run error: %v", err)
	}
	return nil
}
