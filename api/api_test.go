package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gmancoelho/go-bank/models"
	"github.com/gorilla/mux"
)

type MockStorage struct{}

func (m *MockStorage) CreateAccount(account *models.Account) error {
	account.ID = 1
	return nil
}

func (m *MockStorage) DeleteAccount(id int) error {
	return nil
}

func (m *MockStorage) UpdateAccount(account *models.Account) error {
	return nil
}

func (m *MockStorage) GetAccounts() ([]*models.Account, error) {
	return []*models.Account{
		{ID: 1, FirstName: "John", LastName: "Doe", Number: 12345, Balance: 1000},
	}, nil
}

func (m *MockStorage) GetAccountByID(id int) (*models.Account, error) {
	return &models.Account{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Number:    12345,
		Balance:   1000,
		CreatedAt: time.Now(),
	}, nil
}

func TestAPIServer_Start(t *testing.T) {
	store := &MockStorage{}
	server := NewAPIServer(":8080", store)

	req := httptest.NewRequest(http.MethodGet, "/account", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(server.handleAccount))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}
}

func TestHandleCreateAccount(t *testing.T) {
	store := &MockStorage{}
	server := NewAPIServer(":8080", store)

	account := models.CreateAccountRequest{
		FirstName: "John",
		LastName:  "Doe",
	}
	body, _ := json.Marshal(account)

	req := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(server.handleAccount))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}

	var createdAccount models.Account
	err := json.Unmarshal(w.Body.Bytes(), &createdAccount)
	if err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if createdAccount.FirstName != account.FirstName || createdAccount.LastName != account.LastName {
		t.Errorf("expected account with name %s %s, got %s %s",
			account.FirstName, account.LastName, createdAccount.FirstName, createdAccount.LastName)
	}
}
