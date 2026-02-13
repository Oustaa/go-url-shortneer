package handlers

import (
	"github.com/oustaa/go-url-shortner/internal/services"
	"gorm.io/gorm"
)

type V1Handlers struct {
	URL *URLHandlers
}

func GetV1Handlers(db *gorm.DB) *V1Handlers {
	return &V1Handlers{
		URL: &URLHandlers{
			db:      db,
			service: services.NewURLService(db),
		},
	}
}
