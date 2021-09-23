package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	apiKeyOpt  = flag.String("apiKey", getApiKeyFromEnv(), "API key")
	titleOpt   = flag.String("title", "", "Notification title")
	contentOpt = flag.String("content", "", "Notification content")
	urlOpt     = flag.String("url", "", "Notification URL")
)

const apiKeyEnvVariableName = "NOTIFY_ME_API_KEY"

func getApiKeyFromEnv() string {
	return os.Getenv(apiKeyEnvVariableName)
}

func validateInput(apiKey string, title string, content string) error {
	if apiKey == "" {
		return errors.New("API key is required")
	}

	if title == "" && content == "" {
		return errors.New("required: title or content")
	}

	return nil
}

func getTime() string {
	return time.Now().Format(time.UnixDate)
}

// If the content is empty, and URL is defined, it copies the URL to the content.
func normalizeInput(content string, url string) (string, string) {
	if content == "" && url != "" {
		return url, url
	}

	return content, url
}

func main() {
	flag.Parse()
	apiKey := *apiKeyOpt
	content := *contentOpt
	title := *titleOpt
	url := *urlOpt

	content, url = normalizeInput(content, url)

	validateErr := validateInput(apiKey, title, content)

	if validateErr != nil {
		fmt.Println(validateErr)
		os.Exit(1)
	}

	fmt.Printf("Notification started at %s\n", getTime())
	fmt.Printf("Sending to: %s\n", apiKey)
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Content: %s\n", content)

	_, err := notify(apiKey, title, content, url)

	fmt.Printf("Finished at %s\n", getTime())

	if err != nil {
		os.Exit(1)
	}
}
