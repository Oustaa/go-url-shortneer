package handlers

import "net/http"

type URLHandlers struct{}

func (uh *URLHandlers) GetUrls(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET urls route"))
}

func (uh *URLHandlers) PostUrls(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST urls route"))
}
