package action

import (
	"fmt"
	"log"
	"os/exec"
)

func runDockerCompose(args []string) {
	fmt.Println("Running...")
	fmt.Printf("docker %v", args)

	cmd := exec.Command("docker", args[:]...)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
