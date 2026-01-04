# Narwhal Orchestrator

MicroVM orchestrator for managing virtual machines in the Narwhal PaaS using Flintlock.

## Setup

```bash
go mod tidy
```

## Usage

```go
import "narwhal/backend/orchestrator"

orch, _ := orchestrator.NewMicroVMOrchestrator("localhost:9090")
defer orch.Close()

// Create a microVM
uid, _ := orch.CreateMicroVM("my-vm", "default", "", "", 2, 2048)

// Delete microVM
orch.DeleteMicroVM(uid)
```

Run program: `go run .`

Format code: `gofmt -s -w .`

## API

**Core:** `NewMicroVMOrchestrator(addr)`, `Close()`, `GetConnection()`

**MicroVMs:** `CreateMicroVM(name, namespace, kernel, rootfs, vcpus, memory)`, `GetMicroVM(uid)`, `ListMicroVMs(namespace)`, `DeleteMicroVM(uid)`

## Requirements

- Flintlock service running
- Go 1.21+