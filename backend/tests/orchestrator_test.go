package main_test

import (
	"strings"
	"testing"
	"time"

	"narwhal/backend/orchestrator"
)

func TestNewOrchestrator(t *testing.T) {
	orch, err := orchestrator.NewOrchestrator()
	if err != nil {
		t.Skipf("Docker not available: %v", err)
	}
	defer orch.Close()

	if orch == nil {
		t.Fatal("Expected orchestrator to be created, got nil")
	}

	if orch.GetClient() == nil {
		t.Fatal("Expected Docker client to be initialized, got nil")
	}
}

func TestRunAndStopContainer(t *testing.T) {
	orch, err := orchestrator.NewOrchestrator()
	if err != nil {
		t.Skipf("Docker not available: %v", err)
	}
	defer orch.Close()

	containerName := "narwhal-test-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
	imageName := "alpine:latest"
	env := []string{"TEST_VAR=test_value"}

	t.Logf("Starting container: %s with image: %s", containerName, imageName)
	containerID, err := orch.RunContainer(imageName, containerName, env)
	if err != nil {
		t.Fatalf("Failed to run container: %v", err)
	}

	if containerID == "" {
		t.Fatal("Expected container ID, got empty string")
	}
	t.Logf("Container started with ID: %s", containerID)

	info, err := orch.GetContainerInfo(containerID)
	if err != nil {
		t.Fatalf("Failed to get container info: %v", err)
	}

	if !info.State.Running {
		t.Fatal("Expected container to be running")
	}

	t.Logf("Stopping container: %s", containerID)
	err = orch.StopContainer(containerID)
	if err != nil {
		t.Fatalf("Failed to stop container: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	info, err = orch.GetContainerInfo(containerID)
	if err != nil {
		t.Fatalf("Failed to get container info after stop: %v", err)
	}

	if info.State.Running {
		t.Fatal("Expected container to be stopped")
	}

	t.Logf("Removing container: %s", containerID)
	err = orch.RemoveContainer(containerID)
	if err != nil {
		t.Logf("Warning: Failed to remove container: %v", err)
	}
}

func TestRunAndDownContainer(t *testing.T) {
	orch, err := orchestrator.NewOrchestrator()
	if err != nil {
		t.Skipf("Docker not available: %v", err)
	}
	defer orch.Close()

	containerName := "narwhal-test-down-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
	imageName := "alpine:latest"
	env := []string{"TEST_VAR=test_value"}

	t.Logf("Starting container: %s", containerName)
	containerID, err := orch.RunContainer(imageName, containerName, env)
	if err != nil {
		t.Fatalf("Failed to run container: %v", err)
	}
	t.Logf("Container started with ID: %s", containerID)

	info, err := orch.GetContainerInfo(containerID)
	if err != nil {
		t.Fatalf("Failed to get container info: %v", err)
	}

	if !info.State.Running {
		t.Fatal("Expected container to be running")
	}

	t.Logf("Stopping and removing container: %s", containerID)
	err = orch.DownContainer(containerID)
	if err != nil {
		t.Fatalf("Failed to down container: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	_, err = orch.GetContainerInfo(containerID)
	if err == nil {
		t.Fatal("Expected error when getting info for removed container")
	}
}

func TestListContainers(t *testing.T) {
	orch, err := orchestrator.NewOrchestrator()
	if err != nil {
		t.Skipf("Docker not available: %v", err)
	}
	defer orch.Close()

	containers, err := orch.ListContainers(true)
	if err != nil {
		t.Fatalf("Failed to list containers: %v", err)
	}

	t.Logf("Found %d containers", len(containers))

	runningContainers, err := orch.ListContainers(false)
	if err != nil {
		t.Fatalf("Failed to list running containers: %v", err)
	}

	t.Logf("Found %d running containers", len(runningContainers))
}

func TestRunContainerWithExistingName(t *testing.T) {
	orch, err := orchestrator.NewOrchestrator()
	if err != nil {
		t.Skipf("Docker not available: %v", err)
	}
	defer orch.Close()

	containerName := "narwhal-test-existing-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
	imageName := "alpine:latest"

	t.Logf("Starting container (first time): %s", containerName)
	containerID1, err := orch.RunContainer(imageName, containerName, nil)
	if err != nil {
		t.Fatalf("Failed to run container first time: %v", err)
	}
	t.Logf("First container ID: %s", containerID1)

	err = orch.StopContainer(containerID1)
	if err != nil {
		t.Fatalf("Failed to stop container: %v", err)
	}

	t.Logf("Starting container (second time with same name): %s", containerName)
	containerID2, err := orch.RunContainer(imageName, containerName, nil)
	if err != nil {
		t.Fatalf("Failed to run container second time: %v", err)
	}
	t.Logf("Second container ID: %s", containerID2)

	if containerID1 == containerID2 {
		t.Fatal("Expected different container IDs")
	}

	err = orch.DownContainer(containerID2)
	if err != nil {
		t.Logf("Warning: Failed to clean up container: %v", err)
	}
}
