package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func getHeadBranch() (string, error) {
	cmd := exec.Command("git", "remote", "show", "origin")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "HEAD branch:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}
	return "", fmt.Errorf("HEAD branch not found")
}
