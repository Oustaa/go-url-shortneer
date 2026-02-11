package models

type UrlStats struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	ShortURL string `json:"short_url"`
}
