package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const apiURL = "https://api.github.com/users/%s/events"

// Event represents a GitHub activity event.
type Event struct {
	Type    string `json:"type"`
	Repo    struct {
		Name string `json:"name"`
	} `json:"repo"`
}

// FetchUserActivity fetches the recent activity for a GitHub user.
func FetchUserActivity(username string) ([]string, error) {
	url := fmt.Sprintf(apiURL, username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to GitHub API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("user not found or API error")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %v", err)
	}

	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	var activity []string
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			activity = append(activity, fmt.Sprintf("Pushed to %s", event.Repo.Name))
		case "IssuesEvent":
			activity = append(activity, fmt.Sprintf("Opened an issue in %s", event.Repo.Name))
		case "WatchEvent":
			activity = append(activity, fmt.Sprintf("Starred %s", event.Repo.Name))
		default:
			activity = append(activity, fmt.Sprintf("Performed %s on %s", event.Type, event.Repo.Name))
		}
	}

	return activity, nil
}
