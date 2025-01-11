package main

import (
	"github-activity-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
