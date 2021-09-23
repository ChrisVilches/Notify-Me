package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInputCorrect(t *testing.T) {
	var tests = []struct {
		a, b, c string
	}{
		{"x", "x", "x"},
		{"x", "x", ""},
		{"x", "", "x"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("Input: %s %s %s", test.a, test.b, test.c)
		t.Run(testname, func(t *testing.T) {
			assert.Nil(t, validateInput(test.a, test.b, test.c))
		})
	}
}

func TestValidateInputIncorrect(t *testing.T) {
	var tests = []struct {
		a, b, c string
	}{
		{"x", "", ""},
		{"", "x", "x"},
		{"", "x", ""},
		{"", "", "x"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("Input: %s %s %s", test.a, test.b, test.c)
		t.Run(testname, func(t *testing.T) {
			assert.NotNil(t, validateInput(test.a, test.b, test.c))
		})
	}
}

func TestNormalizeInput(t *testing.T) {
	emptyContent := ""
	presentContent := "xxx"
	emptyUrl := ""
	presentUrl := "http"

	var resultContent string
	var resultUrl string

	resultContent, resultUrl = normalizeInput(emptyContent, emptyUrl)
	assert.Equal(t, resultContent, "")
	assert.Equal(t, resultUrl, "")
	resultContent, resultUrl = normalizeInput(emptyContent, presentUrl)
	assert.Equal(t, resultContent, "http")
	assert.Equal(t, resultUrl, "http")

	resultContent, resultUrl = normalizeInput(presentContent, emptyUrl)
	assert.Equal(t, resultContent, "xxx")
	assert.Equal(t, resultUrl, "")
	resultContent, resultUrl = normalizeInput(presentContent, presentUrl)
	assert.Equal(t, resultContent, "xxx")
	assert.Equal(t, resultUrl, "http")
}

func TestCreatePushUrl(t *testing.T) {
	assert.Equal(t, createPushUrl("1", "2", "3", ""), pushServiceUrlBase+"?k=1&t=2&c=3")
	assert.Equal(t, createPushUrl("1", "2", "3", "4"), pushServiceUrlBase+"?k=1&t=2&c=3&u=4")
}
