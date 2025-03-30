package main

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
