package services

import (
	"github.com/oustaa/go-url-shortner/internal/models"
	"gorm.io/gorm"
)

type URLServices struct {
	db *gorm.DB
}

// NewURLServices creates a new URLServices instance
func NewURLServices(db *gorm.DB) *URLServices {
	return &URLServices{
		db: db,
	}
}

// SetDB sets the database connection
func (us *URLServices) SetDB(db *gorm.DB) {
	us.db = db
}

func (us *URLServices) GetUrls() (*[]models.URL, error) {
	var urls []models.URL

	result := us.db.Preload("User").Find(&urls)

	if result.Error != nil {
		return nil, result.Error
	}

	return &urls, nil
}

func (us *URLServices) PostUrls() {
}
