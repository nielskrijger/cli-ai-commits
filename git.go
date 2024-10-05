package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

var execCommand = exec.Command

func isGitRepo() error {
	cmd := execCommand("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return errors.New("current directory is not a git repository")
	}
	return nil
}

func getStagedDiff() (string, error) {
	cmd := execCommand("git", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func formatDiff(diff string) string {
	lines := strings.Split(diff, "\n")

	var formattedLines []string

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "diff") || strings.HasPrefix(line, "index") {
			continue
		}
		formattedLines = append(formattedLines, strings.TrimSpace(line))
	}

	return strings.Join(formattedLines, " ")
}
