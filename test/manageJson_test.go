package internal

import (
	"testing"

	"github.com/arfeloreed/password-gen/internal"
)

func TestCreateJson(t *testing.T) {
	email := "test"
	password := "test"
	message, err := internal.CreateJson(email, password)
	if err != nil {
		t.Fatal(err)
	}
	if message != "success" {
		t.Fatalf("Expected message to be 'success', got %s", message)
	}
}

func TestSearchJson(t *testing.T) {
	email := "test"
	data, err := internal.SearchJson(email)
	if err != nil {
		t.Fatal(err)
	}
	if data.Email != email {
		t.Fatalf("Expected email to be %s, got %s", email, data.Email)
	}
}

func TestEditJson(t *testing.T) {
	email := "test"
	password := "edittest"
	message, err := internal.EditJson(email, password)
	if err != nil {
		t.Fatal(err)
	}
	if message != "success" {
		t.Fatalf("Expected message to be 'success', got %s", message)
	}
}

func TestDeleteEmail(t *testing.T) {
	email := "test"
	message, err := internal.DeleteEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	if message != "success" {
		t.Fatalf("Expected message to be 'success', got %s", message)
	}
}
