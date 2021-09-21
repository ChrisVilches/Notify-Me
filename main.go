package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var (
	apiKeyOpt  = flag.String("apiKey", getApiKeyFromEnv(), "API key")
	titleOpt   = flag.String("title", "", "Notification title")
	contentOpt = flag.String("content", "", "Notification content")
)

// TODO: Put functios in a different file
// and in this file, only CLI related stuff.
// Write a test if possible (that checks INPUT and checks errors are raised when data is missing)

const apiKeyEnvVariableName = "NOTIFY_ME_API_KEY"

const pushServiceUrlBase = "http://xdroid.net/api/message"

func createPushUrl(apiKey string, title string, content string) string {
	title = url.QueryEscape(title)
	content = url.QueryEscape(content)
	return fmt.Sprintf("%s?k=%s&t=%s&c=%s", pushServiceUrlBase, apiKey, title, content)
}

func getApiKeyFromEnv() string {
	return os.Getenv(apiKeyEnvVariableName)
}

// TODO: This method dies. It should return error so that it's handled by the caller.
// Should re-return the return of http.Get ??? うるさい lol
func notify(apiKey string, title string, content string) {
	pushUrl := createPushUrl(apiKey, title, content)
	res, err := http.Get(pushUrl)

	if err != nil {
		fmt.Println("There was an error:")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Status)
}

func validateInput(apiKey string, title string, content string) {
	if apiKey == "" {
		// TODO: can be set using ENV or opts, so change message.
		fmt.Printf("Environment variable %s must be set with the API key.\n", apiKeyEnvVariableName)
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

	notify(apiKey, title, content)
}
