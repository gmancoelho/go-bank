package models

import (
	"encoding/json"
	"testing"
)

func TestApiErrorSerialization(t *testing.T) {
	apiError := ApiError{
		Code:    404,
		Message: "Resource not found",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(apiError)
	if err != nil {
		t.Fatalf("failed to serialize ApiError: %v", err)
	}

	expectedJSON := `{"code":404,"message":"Resource not found"}`
	if string(jsonData) != expectedJSON {
		t.Errorf("expected JSON %s, got %s", expectedJSON, string(jsonData))
	}

	// Deserialize from JSON
	var deserializedError ApiError
	err = json.Unmarshal(jsonData, &deserializedError)
	if err != nil {
		t.Fatalf("failed to deserialize ApiError: %v", err)
	}

	if deserializedError.Code != apiError.Code {
		t.Errorf("expected Code %d, got %d", apiError.Code, deserializedError.Code)
	}

	if deserializedError.Message != apiError.Message {
		t.Errorf("expected Message %s, got %s", apiError.Message, deserializedError.Message)
	}
}
