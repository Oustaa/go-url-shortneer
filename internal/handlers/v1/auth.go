package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ah AuthHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var body CreateAccountRequest
	json.NewDecoder(r.Body).Decode(&body)

	fmt.Printf("%#v", body)

	w.Write([]byte("created successfully"))
}
