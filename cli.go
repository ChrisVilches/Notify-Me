package main

import (
	"errors"
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

func validateInput(apiKey string, title string, content string) error {
	if apiKey == "" {
		return errors.New("API key is required")
	}

	if title == "" || content == "" {
		return errors.New("title and content are required")
	}

	return nil
}

func main() {
	flag.Parse()
	apiKey := *apiKeyOpt
	content := *contentOpt
	title := *titleOpt

	validateErr := validateInput(apiKey, title, content)

	if validateErr != nil {
		fmt.Println(validateErr)
		os.Exit(1)
	}

	fmt.Printf("Sending to: %s\n", apiKey)
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Content: %s\n", content)

	_, err := notify(apiKey, title, content)

	if err != nil {
		os.Exit(1)
	}
}
