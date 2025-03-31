package models

import (
	"testing"
	"time"
)

func TestNewAccount(t *testing.T) {
	firstName := "John"
	lastName := "Doe"

	account := NewAccount(firstName, lastName)

	if account.FirstName != firstName {
		t.Errorf("expected FirstName to be %s, got %s", firstName, account.FirstName)
	}

	if account.LastName != lastName {
		t.Errorf("expected LastName to be %s, got %s", lastName, account.LastName)
	}

	if account.Number <= 0 {
		t.Errorf("expected Number to be greater than 0, got %d", account.Number)
	}

	if time.Since(account.CreatedAt) > time.Second {
		t.Errorf("expected CreatedAt to be recent, got %s", account.CreatedAt)
	}
}

func TestGetID(t *testing.T) {
	account := &Account{ID: 123}

	if account.GetID() != 123 {
		t.Errorf("expected ID to be 123, got %d", account.GetID())
	}
}

func TestGetName(t *testing.T) {
	account := &Account{FirstName: "John", LastName: "Doe"}

	expectedName := "John Doe"
	if account.GetName() != expectedName {
		t.Errorf("expected Name to be %s, got %s", expectedName, account.GetName())
	}
}
