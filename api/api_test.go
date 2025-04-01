package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gmancoelho/go-bank/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandleAccount(t *testing.T) {
	store := NewMockStorage()
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
	store := NewMockStorage()
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

func TestHandleGetAccountByIDNotFound(t *testing.T) {
	store := NewMockStorage()
	server := NewAPIServer(":8080", store)

	req := httptest.NewRequest(http.MethodGet, "/account/999", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleGetAccountById))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var apiError map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, "account not found", apiError["message"])
}

func TestHandleCreateAccount(t *testing.T) {
	store := NewMockStorage()
	server := NewAPIServer(":8080", store)

	payload := `{"firstName": "Jane", "lastName": "Doe"}`
	req := httptest.NewRequest(http.MethodPost, "/account", http.NoBody)
	req.Body = io.NopCloser(bytes.NewReader([]byte(payload)))
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(server.handleAccount))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var account models.Account
	err := json.Unmarshal(w.Body.Bytes(), &account)
	assert.NoError(t, err)
	assert.Equal(t, "Jane", account.FirstName)
	assert.Equal(t, "Doe", account.LastName)
}

func TestHandleDeleteAccount(t *testing.T) {
	store := NewMockStorage()
	server := NewAPIServer(":8080", store)

	req := httptest.NewRequest(http.MethodDelete, "/account/1", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleDeleteAccount))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestHandleDeleteAccountNotFound(t *testing.T) {
	store := NewMockStorage()
	server := NewAPIServer(":8080", store)

	req := httptest.NewRequest(http.MethodDelete, "/account/999", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleDeleteAccount))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var apiError map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, "failed to delete account", apiError["message"])
}

func TestHandleUpdateAccount(t *testing.T) {
	store := NewMockStorage()
	server := NewAPIServer(":8080", store)

	payload := `{"firstName": "Jane", "lastName": "Smith"}`
	req := httptest.NewRequest(http.MethodPut, "/account/1", http.NoBody)
	req.Body = io.NopCloser(bytes.NewReader([]byte(payload)))
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleUpdateAccount))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var account models.Account
	err := json.Unmarshal(w.Body.Bytes(), &account)
	assert.NoError(t, err)
	assert.Equal(t, "Jane", account.FirstName)
	assert.Equal(t, "Smith", account.LastName)
}

func TestHandleUpdateAccountNotFound(t *testing.T) {
	store := NewMockStorage()
	server := NewAPIServer(":8080", store)

	payload := `{"firstName": "Jane", "lastName": "Smith"}`
	req := httptest.NewRequest(http.MethodPut, "/account/999", http.NoBody)
	req.Body = io.NopCloser(bytes.NewReader([]byte(payload)))
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(server.handleUpdateAccount))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var apiError map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, "failed to retrieve account", apiError["message"])
}
