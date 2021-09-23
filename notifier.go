package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const pushServiceUrlBase = "http://xdroid.net/api/message"

// Generates the URL used for the HTTP GET request that triggers the notification.
func createPushUrl(apiKey string, title string, content string, link string) string {
	title = url.QueryEscape(title)
	content = url.QueryEscape(content)
	link = url.QueryEscape(link)

	if link == "" {
		return fmt.Sprintf("%s?k=%s&t=%s&c=%s", pushServiceUrlBase, apiKey, title, content)
	} else {
		return fmt.Sprintf("%s?k=%s&t=%s&c=%s&u=%s", pushServiceUrlBase, apiKey, title, content, link)
	}
}

// Dispatch the notification. Executes a synchronous HTTP request.
func notify(apiKey string, title string, content string, link string) (*http.Response, error) {
	pushUrl := createPushUrl(apiKey, title, content, link)
	res, err := http.Get(pushUrl)
	fmt.Println(res.Status)

	if err != nil {
		fmt.Println("There was an error:")
		fmt.Println(err)
	}

	return res, err
}
