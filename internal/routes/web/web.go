package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func GetWEBRoutes(r chi.Router, db *gorm.DB) {
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, Welcome Home.</h1>"))
	})
}
