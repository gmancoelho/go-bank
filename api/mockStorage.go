package api

import (
	"fmt"

	"github.com/gmancoelho/go-bank/models"
)

type MockStorage struct {
	accounts map[int]*models.Account // In-memory storage for accounts
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		accounts: map[int]*models.Account{
			1: {ID: 1, FirstName: "John", LastName: "Doe", Number: 12345, Balance: 1000},
		},
	}
}

func (m *MockStorage) CreateAccount(account *models.Account) error {
	account.ID = len(m.accounts) + 1
	m.accounts[account.ID] = account
	return nil
}

func (m *MockStorage) GetAccountByID(id int) (*models.Account, error) {
	account, exists := m.accounts[id]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}
	return account, nil
}

func (m *MockStorage) GetAccounts() ([]*models.Account, error) {
	accounts := []*models.Account{}
	for _, account := range m.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (m *MockStorage) DeleteAccount(id int) error {
	if _, exists := m.accounts[id]; !exists {
		return fmt.Errorf("account not found")
	}
	delete(m.accounts, id)
	return nil
}

func (m *MockStorage) UpdateAccount(account *models.Account) error {
	if _, exists := m.accounts[account.ID]; !exists {
		return fmt.Errorf("account not found")
	}
	m.accounts[account.ID] = account
	return nil
}
