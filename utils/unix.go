package utils

import (
	"log"
	"os/exec"
)

func Bash(first string, args ...string) string {
	cmd := exec.Command(first, args...)
	resp, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(resp)
}
