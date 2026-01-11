package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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

	// CLI Commands
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createName := createCmd.String("name", "my-vm", "Name of the MicroVM")
	createCPUs := createCmd.Int("cpus", 2, "Number of vCPUs")
	createMem := createCmd.Int("mem", 2048, "Memory in MB")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.String("id", "", "MicroVM UID to delete")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'create', 'list', or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		fmt.Printf("Creating MicroVM '%s' (cpus: %d, mem: %d)...\n", *createName, *createCPUs, *createMem)
		uid, err := orch.CreateMicroVM(*createName, "default", "", "", int32(*createCPUs), int32(*createMem))
		if err != nil {
			log.Fatalf("Failed to create MicroVM: %v", err)
		}
		fmt.Printf("MicroVM created successfully! UID: %s\n", uid)

	case "list":
		listCmd.Parse(os.Args[2:])
		vms, err := orch.ListMicroVMs("default")
		if err != nil {
			log.Fatalf("Failed to list MicroVMs: %v", err)
		}
		fmt.Printf("Found %d MicroVMs:\n", len(vms))
		for _, vm := range vms {
			// Handle potential nil pointers for safety
			id := "unknown"
			if vm.Spec != nil && vm.Spec.Id != "" {
				id = vm.Spec.Id
			}
			uid := "unknown" 
			if vm.Spec != nil && vm.Spec.Uid != nil {
				uid = *vm.Spec.Uid
			}
			fmt.Printf("- Name: %s | UID: %s | State: %s\n", id, uid, vm.Status.State)
		}

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteID == "" {
			fmt.Println("Error: -id is required for delete command")
			deleteCmd.PrintDefaults()
			os.Exit(1)
		}
		fmt.Printf("Deleting MicroVM %s...\n", *deleteID)
		err := orch.DeleteMicroVM(*deleteID)
		if err != nil {
			log.Fatalf("Failed to delete MicroVM: %v", err)
		}
		fmt.Println("MicroVM deleted successfully.")

	default:
		fmt.Println("expected 'create', 'list', or 'delete' subcommands")
		os.Exit(1)
	}
}
