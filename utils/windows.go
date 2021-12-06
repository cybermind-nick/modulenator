package utils

import (
	"log"
	"os/exec"
)

// Windows utility functions are defined here

// Powershell() is a utility function for setting the path
// to the powerhshell executable

func Powershell(args ...string) string {
	ps, _ := exec.LookPath("powershell.exe")

	cmd := exec.Command(ps, args...)

	resp, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Failed to execute powershell command")
	}

	return string(resp)

}
