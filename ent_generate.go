// +build ignore

package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Run go mod tidy first
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("go mod tidy failed: %v", err)
	}

	// Run ent generation
	cmd = exec.Command("go", "run", "entgo.io/ent/cmd/ent", "generate", "./ent/schema")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("ent generate failed: %v", err)
	}
}