package handlers

import handlers_v1 "github.com/oustaa/go-url-shortner/internal/handlers/v1"

type Handler struct {
	V1 *handlers_v1.URLHandlers
}

func GetHandlers() *Handler {
	return &Handler{
		V1: &handlers_v1.URLHandlers{},
	}
}
