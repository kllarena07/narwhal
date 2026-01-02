# Narwhal Orchestrator (Go)

A Docker orchestrator written in Go for managing containers in the Narwhal PaaS.

## Setup

1. Install Go (if not already installed): https://go.dev/doc/install

2. Install dependencies:
```bash
cd backend
go mod tidy
```

## Usage

### As a Library

```go
package main

import (
    "log"
    "fmt"
    "narwhal/orchestrator"
)

func main() {
    orch, err := orchestrator.NewOrchestrator()
    if err != nil {
        log.Fatal(err)
    }
    defer orch.Close()

    // List containers
    containers, err := orch.ListContainers(true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d containers\n", len(containers))

    // Run a container
    containerID, err := orch.RunContainer(
        "nginx:latest",
        "my-nginx",
        []string{"ENV_VAR=value"},
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Started container: %s\n", containerID)
}
```

### As a Standalone Program

```bash
go run .
```

or

```bash
go run main.go orchestrator.go
```

### Running Tests

Tests are located in the `tests/` directory:

```bash
go test ./tests/
```

## API

- `NewOrchestrator()` - Creates a new orchestrator instance
- `ListContainers(all bool)` - Lists all containers
- `RunContainer(image, name, env)` - Creates and starts a container
- `StopContainer(containerID)` - Stops a running container
- `RemoveContainer(containerID)` - Removes a container
- `GetContainerInfo(containerID)` - Gets detailed info about a container
- `Close()` - Closes the Docker client connection

## Requirements

- Docker daemon must be running
- Go 1.21 or later
