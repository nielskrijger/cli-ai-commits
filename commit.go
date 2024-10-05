package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateMessage(apiKey string) {
	for {
		commitMessage := gitDiff(apiKey)

		fmt.Printf("Suggested Commit Message: %s\n", commitMessage)

		fmt.Print("Are you satisfied with this commit message? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(strings.ToLower(userInput))

		if userInput == "yes" {
			if err := CommitMessage(commitMessage); err != nil {
				log.Fatalf("Error executing git commit: %s", err)
			}
			fmt.Println("Commit successfully created!")
			break
		} else {
			fmt.Println("Generating a new commit message...")
		}
	}
}

func gitDiff(apiKey string) string {
	err := IsGitRepo()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	diff, err := GetStagedDiff()
	if err != nil {
		log.Fatalf("Error getting staged diff: %s", err)
	}

	formattedDiff := FormatDiff(diff)

	gen := &OpenAIGenerator{}
	commitMsg, err := gen.GenerateCommitMsg(apiKey, "Generate a concise commit message for the following changes: "+formattedDiff)
	if err != nil {
		log.Fatalf("Error generating commit message: %s", err)
	}

	return commitMsg
}
