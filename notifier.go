package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const pushServiceUrlBase = "http://xdroid.net/api/message"

func createPushUrl(apiKey string, title string, content string) string {
	title = url.QueryEscape(title)
	content = url.QueryEscape(content)
	return fmt.Sprintf("%s?k=%s&t=%s&c=%s", pushServiceUrlBase, apiKey, title, content)
}

func notify(apiKey string, title string, content string) (*http.Response, error) {
	pushUrl := createPushUrl(apiKey, title, content)
	res, err := http.Get(pushUrl)
	fmt.Println(res.Status)

	if err != nil {
		fmt.Println("There was an error:")
		fmt.Println(err)
	}

	return res, err
}
