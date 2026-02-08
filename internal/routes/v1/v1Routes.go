package v1_routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-url-shortner/internal/handlers"
)

func GetV1Routes(r chi.Router) {
	h := handlers.GetHandlers()

	r.Route("/urls", func(r chi.Router) {
		r.Get("/", h.V1.GetUrls)
		r.Post("/", h.V1.PostUrls)
	})
}
