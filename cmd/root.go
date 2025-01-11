package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity <username>",
	Short: "Fetch recent GitHub activity for a user",
	Long:  "This CLI fetches and displays recent GitHub activity for a specified user.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: Username is required.")
			fmt.Println("Usage: github-activity <username>")
			os.Exit(1)
		}
		username := args[0]
		fetchGitHubActivity(username)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
