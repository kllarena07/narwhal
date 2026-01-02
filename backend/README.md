# Narwhal Orchestrator

Docker orchestrator for managing containers in the Narwhal PaaS.

## Setup

```bash
go mod tidy
```

## Usage

```go
import "narwhal/orchestrator"

orch, _ := orchestrator.NewOrchestrator()
defer orch.Close()

// Start a container
id, _ := orch.RunContainer("nginx:latest", "my-app", []string{"PORT=3000"})

// Stop and remove
orch.DownContainer(id)
```

Run tests: `go test ./tests/`

Run program: `go run .`

Format code: `gofmt -s -w .`

## API

**Core:** `NewOrchestrator()`, `Close()`, `GetClient()`

**Containers:** `RunContainer(image, name, env)`, `StopContainer(id)`, `RemoveContainer(id)`, `DownContainer(id)`, `ForceRemoveContainer(id)`, `ListContainers(all)`, `GetContainerInfo(id)`

## Requirements

- Docker daemon running
- Go 1.21+
