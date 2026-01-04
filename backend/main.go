package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"narwhal/backend/orchestrator"
)

func main() {
	flintlockAddr := os.Getenv("FLINTLOCK_ADDR")
	if flintlockAddr == "" {
		flintlockAddr = "localhost:9090"
	}

	orch, err := orchestrator.NewMicroVMOrchestrator(flintlockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer orch.Close()

	fmt.Printf("Connected to MicroVM Orchestrator at %s\n", flintlockAddr)
}
