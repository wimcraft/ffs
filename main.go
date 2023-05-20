package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("bash", "-c", "cat ~/.bash_history | tail -n 1\n")
	originalCommand, _ := cmd.CombinedOutput()
	output, err := exec.Command("bash", "-c", string(originalCommand)).CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		// Check if the error message matches the one we want to handle
		if strings.Contains(string(output), "Could not open lock file") {
			// If the error matches, create the corrected command
			fixedCmd := "sudo " + string(originalCommand)

			fmt.Printf("Fixed Command:\n%s\n", fixedCmd)
		} else {
			// Print the original error message if it doesn't match the one we want to handle
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	}
}
