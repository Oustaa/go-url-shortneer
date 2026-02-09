package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type URLHandlers struct {
	db *gorm.DB
}

func (uh *URLHandlers) GetUrls(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1, GET urls route"))
}

func (uh *URLHandlers) PostUrls(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1, POST urls route"))
}
