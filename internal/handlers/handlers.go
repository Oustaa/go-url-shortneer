package handlers

import (
	handlers_v1 "github.com/oustaa/go-url-shortner/internal/handlers/v1"
	"gorm.io/gorm"
)

type Handler struct {
	V1 *handlers_v1.V1Handlers
}

func GetHandlers(db *gorm.DB) *Handler {
	return &Handler{
		V1: handlers_v1.GetV1Handlers(db),
	}
}
