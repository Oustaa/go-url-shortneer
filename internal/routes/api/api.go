package api

import (
	"github.com/go-chi/chi/v5"
	routes_v1 "github.com/oustaa/go-url-shortner/internal/routes/api/v1"
	"gorm.io/gorm"
)

func GetAPIRoutes(r chi.Router, db *gorm.DB) {
	r.Route("/v1", func(r chi.Router) {
		routes_v1.GetV1Routes(r, db)
	})
}
