package routes

import (
	"github.com/go-chi/chi/v5"
	api_routes "github.com/oustaa/go-url-shortner/internal/routes/api"
	web_routes "github.com/oustaa/go-url-shortner/internal/routes/web"
	"gorm.io/gorm"
)

func GetRoutes(db *gorm.DB) *chi.Mux {
	router := chi.NewMux()

	router.Route("/api", func(r chi.Router) {
		api_routes.GetAPIRoutes(r, db)
	})

	web_routes.GetWEBRoutes(router, db)

	return router
}
