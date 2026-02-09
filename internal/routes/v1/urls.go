package v1_routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-url-shortner/internal/handlers"
	"gorm.io/gorm"
)

func GetUrlsRoutes(r chi.Router, db *gorm.DB) {
	h := handlers.GetHandlers(db)

	r.Route("/urls", func(r chi.Router) {
		r.Get("/", h.V1.URL.GetUrls)
		r.Post("/", h.V1.URL.PostUrls)
	})
}
