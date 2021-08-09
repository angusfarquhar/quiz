package main

import "testing"

func TestMain(t *testing.T) {
	result := "Hello Angus!"

	if result != "Hello World!" {
		t.Errorf("TestMain failed! Expected: %v, got %v", result, "Hello World!")
	}
}
