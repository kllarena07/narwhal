package main

import (
	"fmt"
	"log"

	"narwhal/orchestrator"
)

func main() {
	orch, err := orchestrator.NewOrchestrator()
	if err != nil {
		log.Fatal(err)
	}
	defer orch.Close()

	fmt.Println("Orchestrator initialized successfully")
	fmt.Printf("Docker client: %v\n", orch.GetClient())
}
