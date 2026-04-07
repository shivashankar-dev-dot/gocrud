package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/create", CreateUser).Methods("POST")

	log.Println("Server running on Port 8080")
	http.ListenAndServe(":8080", r)

}
