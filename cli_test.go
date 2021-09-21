package main

import "testing"

// TODO: Are these tests following good practices?

func TestValidateInputIncorrect(t *testing.T) {
	errMsg := "not handled as incorrect"
	if validateInput("", "", "") == nil {
		t.Errorf(errMsg)
	}

	if validateInput("x", "", "") == nil {
		t.Errorf(errMsg)
	}

	if validateInput("", "x", "") == nil {
		t.Errorf(errMsg)
	}

	if validateInput("", "", "x") == nil {
		t.Errorf(errMsg)
	}

	if validateInput("", "", "x") == nil {
		t.Errorf(errMsg)
	}
}

func TestValidateInputCorrect(t *testing.T) {
	if validateInput("x", "x", "x") != nil {
		t.Errorf("input is correct but it's handled as incorrect")
	}
}
