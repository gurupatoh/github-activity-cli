package cmd

import (
	"fmt"
	"github-activity-cli/github"
)

func fetchGitHubActivity(username string) {
	events, err := github.FetchUserActivity(username)
	if err != nil {
		fmt.Printf("Error fetching activity for user '%s': %v\n", username, err)
		return
	}

	if len(events) == 0 {
		fmt.Printf("No recent activity found for user '%s'.\n", username)
		return
	}

	fmt.Printf("Recent activity for '%s':\n", username)
	for _, event := range events {
		fmt.Println(event)
	}
}
