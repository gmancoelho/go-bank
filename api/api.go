package api

import (
	"log"
	"net/http"

	m "github.com/gmancoelho/go-bank/models"
	repo "github.com/gmancoelho/go-bank/repository"
	u "github.com/gmancoelho/go-bank/utils"
	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			u.WriteJSON(w, http.StatusInternalServerError,
				m.ApiError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
		}
	}
}

type APIServer struct {
	address string
	store   repo.Storage
}

func NewAPIServer(address string, store repo.Storage) *APIServer {
	return &APIServer{address: address, store: store}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Start() error {
	router := s.setupRouter()

	log.Printf("JSON API server started on %s\n", s.address)

	return http.ListenAndServe(s.address, router)
}

func (s *APIServer) setupRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(LoggingMiddleware)

	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetAccountById))

	return router
}
