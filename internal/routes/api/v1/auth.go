package v1_routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-url-shortner/internal/handlers"
	"gorm.io/gorm"
)

func GetAuthRoutes(r chi.Router, db *gorm.DB) {
	h := handlers.GetHandlers(db)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/create-account", h.V1.Auth.SignIn)
	})
}
