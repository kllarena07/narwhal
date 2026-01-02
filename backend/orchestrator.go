package orchestrator

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// Orchestrator manages Docker containers and operations
type Orchestrator struct {
	client *client.Client
	ctx    context.Context
}

// NewOrchestrator creates a new orchestrator instance with a Docker client
func NewOrchestrator() (*Orchestrator, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}

	// Verify Docker connection
	ctx := context.Background()
	_, err = cli.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Docker daemon: %w", err)
	}

	return &Orchestrator{
		client: cli,
		ctx:    ctx,
	}, nil
}

// GetClient returns the underlying Docker client
func (o *Orchestrator) GetClient() *client.Client {
	return o.client
}

// ListContainers returns a list of all containers
func (o *Orchestrator) ListContainers(all bool) ([]types.Container, error) {
	containers, err := o.client.ContainerList(o.ctx, container.ListOptions{
		All: all,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}
	return containers, nil
}

// RunContainer runs a container with the specified configuration
func (o *Orchestrator) RunContainer(imageName string, name string, env []string) (string, error) {
	// Pull the image if it doesn't exist locally
	out, err := o.client.ImagePull(o.ctx, imageName, image.PullOptions{})
	if err == nil {
		// Consume the pull output
		io.Copy(io.Discard, out)
		out.Close()
	}
	// Note: We continue even if image pull fails, as the image might already exist locally

	// Create container
	resp, err := o.client.ContainerCreate(
		o.ctx,
		&container.Config{
			Image: imageName,
			Env:   env,
		},
		&container.HostConfig{},
		nil,
		nil,
		name,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %w", err)
	}

	// Start container
	if err := o.client.ContainerStart(o.ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start container: %w", err)
	}

	return resp.ID, nil
}

// StopContainer stops a running container
func (o *Orchestrator) StopContainer(containerID string) error {
	timeout := 10 // seconds
	return o.client.ContainerStop(o.ctx, containerID, container.StopOptions{
		Timeout: &timeout,
	})
}

// RemoveContainer removes a container
func (o *Orchestrator) RemoveContainer(containerID string) error {
	return o.client.ContainerRemove(o.ctx, containerID, container.RemoveOptions{})
}

// GetContainerInfo returns information about a specific container
func (o *Orchestrator) GetContainerInfo(containerID string) (types.ContainerJSON, error) {
	info, err := o.client.ContainerInspect(o.ctx, containerID)
	if err != nil {
		return types.ContainerJSON{}, fmt.Errorf("failed to inspect container: %w", err)
	}
	return info, nil
}

// Close closes the Docker client connection
func (o *Orchestrator) Close() error {
	return o.client.Close()
}
