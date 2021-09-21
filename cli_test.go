package main

import (
	"fmt"
	"testing"
)

func TestValidateInputIncorrect(t *testing.T) {
	var tests = []struct {
		a, b, c string
	}{
		{"x", "x", ""},
		{"", "x", "x"},
		{"x", "", "x"},
		{"", "", ""},
		{"", "x", ""},
		{"", "", "x"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("Input: %s %s %s", test.a, test.b, test.c)
		t.Run(testname, func(t *testing.T) {
			if validateInput(test.a, test.b, test.c) == nil {
				t.Errorf("input is incorrect, but was handled as correct")
			}
		})
	}
}

func TestValidateInputCorrect(t *testing.T) {
	if validateInput("x", "x", "x") != nil {
		t.Errorf("input is correct but it's handled as incorrect")
	}
}
