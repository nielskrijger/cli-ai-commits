package main

import (
	"fmt"
	"log"
)

func GenerateMessage(apiKey string) {
	commitMessage := gitDiff(apiKey)

	fmt.Printf("%s\n", commitMessage)
}

func gitDiff(apiKey string) string {
	err := isGitRepo()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	diff, err := getStagedDiff()
	if err != nil {
		log.Fatalf("Error getting staged diff: %s", err)
	}

	formattedDiff := formatDiff(diff)

	gen := &OpenAIGenerator{}
	commitMsg, err := gen.GenerateCommitMsg(apiKey, "Generate a concise commit message for the following changes: "+formattedDiff)
	if err != nil {
		log.Fatalf("Error generating commit message: %s", err)
	}

	return commitMsg
}
