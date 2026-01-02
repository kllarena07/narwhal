package main_test

import (
	"testing"

	"narwhal/orchestrator"
)

func TestNewOrchestrator(t *testing.T) {
	// Skip if Docker is not available
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
