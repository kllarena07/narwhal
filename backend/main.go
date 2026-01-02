package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"narwhal/orchestrator"
)

func main() {
	orch, err := orchestrator.NewOrchestrator()
	if err != nil {
		log.Fatal(err)
	}
	defer orch.Close()

	containerName := "narwhal-app"
	imageName := "nginx:alpine"

	fmt.Printf("Starting container %s with image %s...\n", containerName, imageName)
	containerID, err := orch.RunContainer(imageName, containerName, nil)
	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}

	fmt.Printf("Container started successfully! ID: %s\n", containerID)
	fmt.Println("Press Ctrl+C to stop the container...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nStopping container...")
	err = orch.DownContainer(containerID)
	if err != nil {
		log.Fatalf("Failed to stop container: %v", err)
	}

	fmt.Println("Container stopped and removed successfully!")
}
