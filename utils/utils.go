package utils

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, statusCode int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
