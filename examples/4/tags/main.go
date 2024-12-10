package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var user User

type User struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func getToken(user User) string {
	// Pretty dumb token generation, just for the sake of example
	return "token_" + user.Username
}

type Response struct {
	Token string `json:"token"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token := getToken(user)
	response := Response{Token: token}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("POST /login", HandleLogin)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
