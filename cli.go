package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	apiKeyOpt  = flag.String("apiKey", getApiKeyFromEnv(), "API key")
	titleOpt   = flag.String("title", "", "Notification title")
	contentOpt = flag.String("content", "", "Notification content")
)

const apiKeyEnvVariableName = "NOTIFY_ME_API_KEY"

func getApiKeyFromEnv() string {
	return os.Getenv(apiKeyEnvVariableName)
}

func validateInput(apiKey string, title string, content string) {
	if apiKey == "" {
		fmt.Println("API key is required")
		os.Exit(1)
	}

	if title == "" || content == "" {
		fmt.Println("Title and content are required")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	apiKey := *apiKeyOpt
	content := *contentOpt
	title := *titleOpt

	validateInput(apiKey, title, content)

	fmt.Printf("Sending to: %s\n", apiKey)
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Content: %s\n", content)

	_, err := notify(apiKey, title, content)

	if err != nil {
		os.Exit(1)
	}
}
