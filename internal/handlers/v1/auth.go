package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oustaa/go-url-shortner/internal/services"
	"github.com/oustaa/go-url-shortner/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthHandler struct {
	db      *gorm.DB
	service *services.UserService
}

func (ah AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var body services.UserPayload
	json.NewDecoder(r.Body).Decode(&body)

	user, err := ah.service.CreateUser(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(struct {
		JWT string `json:"jwt"`
	}{
		JWT: token,
	})
}

func (ah AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var body LoginPayload
	json.NewDecoder(r.Body).Decode(&body)

	err := bcrypt.CompareHashAndPassword([]byte(""), []byte(body.Password))
	if err != nil {
		fmt.Println("Password does not match:", err)
		return
	}
}
