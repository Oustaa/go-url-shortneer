package services

import (
	"github.com/oustaa/go-url-shortner/internal/models"
	"gorm.io/gorm"
)

type URLService struct {
	db *gorm.DB
}

func NewURLService(db *gorm.DB) *URLService {
	return &URLService{
		db: db,
	}
}

func (us *URLService) SetDB(db *gorm.DB) {
	us.db = db
}

func (us *URLService) GetUserUrls(userID int64) (*[]models.URL, error) {
	var urls []models.URL

	result := us.db.Where("user_id", userID).Find(&urls)

	if result.Error != nil {
		return nil, result.Error
	}

	return &urls, nil
}

func (us *URLService) GetUrls() (*[]models.URL, error) {
	var urls []models.URL

	result := us.db.Preload("User").Find(&urls)

	if result.Error != nil {
		return nil, result.Error
	}

	return &urls, nil
}

func (us *URLService) PostUrls() {
}
