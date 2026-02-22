package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oustaa/go-url-shortner/internal/services"
	"github.com/oustaa/go-url-shortner/internal/utils"
	"gorm.io/gorm"
)

type URLHandlers struct {
	db      *gorm.DB
	service *services.URLService
}

func (uh *URLHandlers) GetUrls(w http.ResponseWriter, r *http.Request) {
	id := utils.GetUserID(r)

	urls, err := uh.service.GetUserUrls(id)
	if err != nil {
		http.Error(w, "Enable to get the urls.", http.StatusInternalServerError)
		return
	}

	utils.SendResponce(w, urls)
}

func (uh *URLHandlers) PostUrls(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserID(r)

	type BodyStruct struct {
		LongURL string `json:"longUrl"`
	}

	var body BodyStruct

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Body invalid, %#v", err.Error()), http.StatusInternalServerError)
		return
	}

	url, err := uh.service.PostUrls(body.LongURL, userID)
	if err != nil {
		http.Error(w, "Enable to create url", http.StatusInternalServerError)
		return
	}

	utils.SendResponce(w, url)
}
