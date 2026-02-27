package services

import (
	"errors"

	"github.com/oustaa/go-url-shortner/internal/models"
	"github.com/oustaa/go-url-shortner/internal/utils"
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

func (us *URLService) PostUrls(longURL string, userID int64) (*models.URL, error) {
	var url models.URL
	var shortURL string

	for {
		shortURL = utils.EncodeURL(longURL)

		err := us.db.
			Where("short_url = ?", shortURL).
			First(&url).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}

			return nil, err
		}
	}

	newURL := models.URL{
		LongURL:  longURL,
		ShortURL: shortURL,
		UserID:   &userID,
		Visits:   0,
	}

	if err := us.db.Create(&newURL).Error; err != nil {
		return nil, err
	}

	return &newURL, nil
}

func (us *URLService) GetURLByShortHash(shortURL string) (*models.URL, error) {
	var url models.URL

	result := us.db.Where("short_url = ?", shortURL).First(&url)

	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}
