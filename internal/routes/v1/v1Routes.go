package v1_routes

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// GetV1Routes registers all v1 API routes
func GetV1Routes(r chi.Router, db *gorm.DB) {
	GetUrlsRoutes(r, db)
}
