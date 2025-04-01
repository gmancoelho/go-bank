package models

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
