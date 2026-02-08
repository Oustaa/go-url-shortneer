package routes

import (
	"github.com/go-chi/chi/v5"
	v1_routes "github.com/oustaa/go-url-shortner/internal/routes/v1"
)

func GetRoutes() *chi.Mux {
	router := chi.NewMux()

	router.Route("/v1", func(r chi.Router) {
		v1_routes.GetV1Routes(r)
	})

	return router
}
