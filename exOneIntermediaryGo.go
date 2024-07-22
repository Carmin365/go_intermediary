package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := helloWorld()
	if result != "Hello, World!" {
		t.Error("Expected 'Hello, World!', but got '%s'", result)
	}
}

func helloWorld() string {
	return "Hello Wold!"
}
