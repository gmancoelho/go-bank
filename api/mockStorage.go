package api

import (
	"fmt"

	"github.com/gmancoelho/go-bank/models"
)

type MockStorage struct{}

func (m *MockStorage) CreateAccount(account *models.Account) error {
	account.ID = 1
	return nil
}

func (m *MockStorage) GetAccountByID(id int) (*models.Account, error) {
	return &models.Account{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Number:    12345,
		Balance:   1000,
	}, nil
}

func (m *MockStorage) GetAccounts() ([]*models.Account, error) {
	return []*models.Account{
		{ID: 1, FirstName: "John", LastName: "Doe", Number: 12345, Balance: 1000},
	}, nil
}

func (m *MockStorage) DeleteAccount(id int) error {
	if id != 1 {
		return fmt.Errorf("account not found")
	}
	return nil
}
