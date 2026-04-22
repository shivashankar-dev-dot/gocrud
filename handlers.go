package main

import (
	"encoding/json"
	"net/http"

	"github.com/lib/pq"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	err := DB.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)

	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {

			if pqErr.Code == "23505" {
				http.Error(w, "Email already Exists", http.StatusConflict)
				return
			}
		}
	}
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	rows, _ := DB.Query("SELECT *from users")

	var users []User

	for rows.Next() {

	}

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			http.Error(w, pqErr, http.StatusConflict)
			return
		}
	}

}
