package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/oustaa/go-url-shortner/internal/services"
	"gorm.io/gorm"
)

type URLHandlers struct {
	db      *gorm.DB
	service *services.URLService
}

func (uh *URLHandlers) GetUrls(w http.ResponseWriter, r *http.Request) {
	urls, err := uh.service.GetUrls()
	if err != nil {
		http.Error(w, "Enable to get the urls.", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(urls)
	if err != nil {
		http.Error(w, "Enable to encode the urls.", http.StatusInternalServerError)
		return
	}
}

func (uh *URLHandlers) PostUrls(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1, POST urls route"))
}
