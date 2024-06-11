package internal

import (
	"io"
	"os"
	"testing"

	"github.com/arfeloreed/password-gen/internal"
)

func TestMode(t *testing.T) {
	// Create a temporary file to simulate user input
	in, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(in.Name())
	defer in.Close()

	// Write the simulated user input to the temporary file
	_, err = io.WriteString(in, "1\n"+"test\n")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}

	// Save the original stdin
	origStdin := os.Stdin
	defer func() { os.Stdin = origStdin }()

	// Redirect stdin to read from the temporary file
	os.Stdin = in

	// Run the Mode function
	mode, email, passLen := internal.Mode()

	// Check the results
	if mode != 1 {
		t.Fatalf("Expected mode to be 1, got %d", mode)
	}
	if email != "test" {
		t.Fatalf("Expected email to be 'test', got %s", email)
	}
	if passLen != 0 {
		t.Fatalf("Expected passLen to be 0, got %d", passLen)
	}
}
