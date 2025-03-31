package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gmancoelho/go-bank/models"
	"github.com/gmancoelho/go-bank/utils"
	r "github.com/gmancoelho/go-bank/repository"
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
	store r.Storage
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{address: address}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleAccount))

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

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("Get account request received")
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("Create account request received")
	vars := mux.Vars(r)
	log.Println(vars)
	return utils.WriteJSON(w, http.StatusCreated, &models.Account{})
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("Delete account request received")
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
