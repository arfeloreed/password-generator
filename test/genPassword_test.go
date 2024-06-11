package internal

import (
	"testing"

	"github.com/arfeloreed/password-gen/internal"
)

func TestGeneratePassword(t *testing.T) {
	var length int = 15
	password := internal.GeneratePassword(length)
	if len(password) != length {
		t.Fatalf("Expected password length to be %d, got %d", length, len(password))
	}
}
