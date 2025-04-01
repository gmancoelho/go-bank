package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiErrorSerialization(t *testing.T) {
	apiError := ApiError{
		Code:    404,
		Message: "Resource not found",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(apiError)
	assert.NoError(t, err, "failed to serialize ApiError")

	expectedJSON := `{"code":404,"message":"Resource not found"}`
	assert.JSONEq(t, expectedJSON, string(jsonData), "serialized JSON does not match expected JSON")

	// Deserialize from JSON
	var deserializedError ApiError
	err = json.Unmarshal(jsonData, &deserializedError)
	assert.NoError(t, err, "failed to deserialize ApiError")

	assert.Equal(t, apiError.Code, deserializedError.Code, "deserialized Code does not match")
	assert.Equal(t, apiError.Message, deserializedError.Message, "deserialized Message does not match")
}
