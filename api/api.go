package api

import (
	"log"
	"net/http"

	m "github.com/gmancoelho/go-bank/models"
	r "github.com/gmancoelho/go-bank/repository"
	"github.com/gmancoelho/go-bank/utils"
	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError,
				m.ApiError{
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
