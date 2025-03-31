package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gmancoelho/go-bank/models"
	r "github.com/gmancoelho/go-bank/repository"
	"github.com/gmancoelho/go-bank/utils"
	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError,
				ApiError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}
	}
}

type APIServer struct {
	address string
	store   r.Storage
}

func NewAPIServer(address string, store r.Storage) *APIServer {
	return &APIServer{address: address, store: store}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetAccountByID))

	log.Println("JSON API server started on", s.address)

	return http.ListenAndServe(s.address, router)
}

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

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	log.Println("Get account request received")
	return utils.WriteJSON(w, http.StatusCreated, &models.Account{})
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	storage := s.store
	accounts, err := storage.GetAccounts()

	if err != nil {
		return nil
	}

	if len(accounts) == 0 {
		return utils.WriteJSON(w, http.StatusOK, []models.Account{})
	}

	return utils.WriteJSON(w, http.StatusCreated, accounts)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	createAccReq := new(models.CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
		return err
	}

	account := models.NewAccount(createAccReq.FirstName, createAccReq.LastName)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusCreated, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("Delete account request received")
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
