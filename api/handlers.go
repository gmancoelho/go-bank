package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "github.com/gmancoelho/go-bank/models"
	u "github.com/gmancoelho/go-bank/utils"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == http.MethodGet {
		return s.handleGetAccount(w, r)
	}

	if r.Method == http.MethodPost {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.getAccountById(w, r)
	}

	if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	storage := s.store
	accounts, err := storage.GetAccounts()

	if err != nil {
		return nil
	}

	if len(accounts) == 0 {
		return u.WriteJSON(w, http.StatusOK, []m.Account{})
	}

	return u.WriteJSON(w, http.StatusCreated, accounts)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	createAccReq := new(m.CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid request payload",
		})
	}

	if createAccReq.FirstName == "" || createAccReq.LastName == "" {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: "first name and last name cannot be empty",
		})
	}

	account := m.NewAccount(createAccReq.FirstName, createAccReq.LastName)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return u.WriteJSON(w, http.StatusCreated, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := parseAccountID(r)
	if err != nil {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := s.store.DeleteAccount(id); err != nil {
		if err.Error() == fmt.Sprintf("account id %d not found", id) {
			return u.WriteJSON(w, http.StatusNotFound, m.ApiError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		}
		return u.WriteJSON(w, http.StatusInternalServerError, m.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "failed to delete account",
		})
	}

	return u.WriteJSON(w, http.StatusNoContent, map[string]int{"deleted": id})
}

func (s *APIServer) getAccountById(w http.ResponseWriter, r *http.Request) error {
	id, err := parseAccountID(r)
	if err != nil {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}

	if account == nil {
		return fmt.Errorf("account not found")
	}

	return u.WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	// Parse the account ID from the URL
	id, err := parseAccountID(r)
	if err != nil {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Decode the request payload
	updateAccReq := new(m.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(updateAccReq); err != nil {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: "invalid request payload",
		})
	}

	// Validate the request payload
	if updateAccReq.FirstName == "" || updateAccReq.LastName == "" {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: "first name and last name cannot be empty",
		})
	}

	// Retrieve the existing account
	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return u.WriteJSON(w, http.StatusInternalServerError, m.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "failed to retrieve account",
		})
	}
	if account == nil {
		return u.WriteJSON(w, http.StatusNotFound, m.ApiError{
			Code:    http.StatusNotFound,
			Message: "account not found",
		})
	}

	// Update the account fields
	account.FirstName = updateAccReq.FirstName
	account.LastName = updateAccReq.LastName

	// Save the updated account
	if err := s.store.UpdateAccount(account); err != nil {
		return u.WriteJSON(w, http.StatusInternalServerError, m.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "failed to update account",
		})
	}

	// Return the updated account
	return u.WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(m.TransferRequest)

	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return u.WriteJSON(w, http.StatusBadRequest, m.ApiError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	defer r.Body.Close()

	return u.WriteJSON(w, http.StatusOK, transferReq)
}

func parseAccountID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid or missing account ID")
	}
	return id, nil
}
