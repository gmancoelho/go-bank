package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	firstName := "John"
	lastName := "Doe"

	account := NewAccount(firstName, lastName)

	assert.Equal(t, firstName, account.FirstName, "expected FirstName to match")
	assert.Equal(t, lastName, account.LastName, "expected LastName to match")
	assert.Greater(t, account.Number, int64(0), "expected Number to be greater than 0")
	assert.WithinDuration(t, time.Now().UTC(), account.CreatedAt, time.Second, "expected CreatedAt to be recent")
}

func TestGetID(t *testing.T) {
	account := &Account{ID: 123}

	assert.Equal(t, 123, account.GetID(), "expected ID to match")
}

func TestGetName(t *testing.T) {
	account := &Account{FirstName: "John", LastName: "Doe"}

	expectedName := "John Doe"
	assert.Equal(t, expectedName, account.GetName(), "expected Name to match")
}
