package main

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func WriteError(w http.ResponseWriter, appError AppError) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(appError.Status)
	json.NewEncoder(w).Encode(appError)
}

var (
	ErrInternal = AppError{
		Code:    "internal_server_error",
		Message: "Internal Server Error",
		Status:  500,
	}

	ErrDB = AppError{
		Code:    "database_error",
		Message: "Internal Server Error",
		Status:  500,
	}
)
