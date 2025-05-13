package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Alice", "Hello, Alice!"},
		{"Bob", "Hello, Bob!"},
		{"", "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Greet(tt.name)
			if actual != tt.expected {
				t.Errorf("Greet(%s) = %s; want %s", tt.name, actual, tt.expected)
			}
		})
	}
}

func TestMainWithUserInput(t *testing.T) {
	// Simulating user input by redirecting stdin to a buffer
	input := "Alice\n"
	expected := "Hello, Alice!\n"

	// Redirect stdin
	oldStdin := os.Stdin
	os.Stdin = bytes.NewBufferString(input)
	defer func() { os.Stdin = oldStdin }() // Restore original stdin after the test

	// Capture output
	var buf bytes.Buffer
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&buf)
	if err != nil {
		t.Errorf("Error reading input: %v", err)
	}

	// Check if the output is what we expected
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}
