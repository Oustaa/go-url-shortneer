package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-url-shortner/internal/services"
	"github.com/oustaa/go-url-shortner/internal/utils"
	"gorm.io/gorm"
)

type URLHandlers struct {
	db      *gorm.DB
	service *services.URLService
}

func (uh *URLHandlers) GetUrls(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserID(r)

	urls, err := uh.service.GetUserUrls(userID)
	if err != nil {
		http.Error(w, "Enable to get the urls.", http.StatusInternalServerError)
		return
	}

	utils.SendResponce(w, urls)
}

func (uh *URLHandlers) GetArchivedUrls(w http.ResponseWriter, r *http.Request) {
	userID := utils.GetUserID(r)

	urls, err := uh.service.GetArchivedUrls(userID)
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

func (uh *URLHandlers) DeleteURL(w http.ResponseWriter, r *http.Request) {
	// check if the url does exists
	urlHash := chi.URLParam(r, "urlHash")

	fmt.Println(urlHash)

	url, err := uh.service.GetURLByShortHash(urlHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Short URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if that url belongs to the connected user
	userID := utils.GetUserID(r)
	if userID != *url.UserID {
		http.Error(w, "Can't delete url", http.StatusForbidden)
		return
	}

	// delete the url
	if err := uh.service.DeleteURL(url.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success response
	w.WriteHeader(http.StatusNoContent)
}

func RedirectURL(service services.URLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortHash := strings.TrimPrefix(r.URL.Path, "/")

		url, err := service.GetURLByShortHash(shortHash)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "Short URL not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// increate visites for the url
		service.IncreaseVisits(url)

		http.Redirect(w, r, url.LongURL, http.StatusPermanentRedirect)
	}
}
