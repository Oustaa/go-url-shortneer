package routes

import (
	"github.com/go-chi/chi/v5"
	routes_v1 "github.com/oustaa/go-url-shortner/internal/routes/v1"
	"gorm.io/gorm"
)

func GetRoutes(db *gorm.DB) *chi.Mux {
	router := chi.NewMux()

	router.Route("/v1", func(r chi.Router) {
		routes_v1.GetV1Routes(r, db)
	})

	return router
}
