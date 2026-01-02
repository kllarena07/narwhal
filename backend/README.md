# Narwhal Orchestrator

Orchestrator for managing Docker containers and microVMs in the Narwhal PaaS.

Supports:
- **Docker containers** - Lightweight container management
- **MicroVMs via Flintlock** - Full VM lifecycle management using Firecracker/Cloud Hypervisor

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

## API

**Core:** `NewOrchestrator()`, `Close()`, `GetClient()`

**Containers:** `RunContainer(image, name, env)`, `StopContainer(id)`, `RemoveContainer(id)`, `DownContainer(id)`, `ForceRemoveContainer(id)`, `ListContainers(all)`, `GetContainerInfo(id)`

**MicroVMs:** `NewMicroVMOrchestrator(addr)`, `CreateMicroVM(name, namespace, kernel, rootfs, vcpus, memoryMB)`, `GetMicroVM(uid)`, `ListMicroVMs(namespace)`, `DeleteMicroVM(uid)`

## MicroVM Images

Flintlock uses container images for kernels and OS volumes. Recommended images:
- **Kernels**: `ghcr.io/liquidmetal-dev/flintlock-kernel:5.10.77` or `:4.19.215`
- **OS**: `ghcr.io/liquidmetal-dev/capmvm-kubernetes:1.23.5` (Ubuntu 20.04 based)

See [Flintlock Image Guide](https://flintlock.liquidmetal.dev/docs/guides/images/) for more options.

## Requirements

- Docker daemon running (for containers)
- Flintlock daemon running (for microVMs) - see [Flintlock docs](https://flintlock.liquidmetal.dev/)
- Go 1.21+
