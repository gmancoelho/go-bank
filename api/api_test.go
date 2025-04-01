package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gmancoelho/go-bank/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandleAccountIfAccountsIsEmpy(t *testing.T) {
	store := &MockStorage{}
	server := NewAPIServer(":8080", store)

	req := httptest.NewRequest(http.MethodGet, "/account", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(server.handleAccount))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var accounts []*models.Account
	err := json.Unmarshal(w.Body.Bytes(), &accounts)
	assert.NoError(t, err)
	assert.Len(t, accounts, 1)
	assert.Equal(t, "John", accounts[0].FirstName)
}

func TestHandleGetAccountByID(t *testing.T) {
	store := &MockStorage{}
	server := NewAPIServer(":8080", store)

	req := httptest.NewRequest(http.MethodGet, "/account/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleGetAccountById))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var account models.Account
	err := json.Unmarshal(w.Body.Bytes(), &account)
	assert.NoError(t, err)
	assert.Equal(t, "John", account.FirstName)
	assert.Equal(t, int64(12345), account.Number)
}

func TestHandleDeleteAccount(t *testing.T) {
	store := &MockStorage{}
	server := NewAPIServer(":8080", store)

	req := httptest.NewRequest(http.MethodDelete, "/account/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleDeleteAccount))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
