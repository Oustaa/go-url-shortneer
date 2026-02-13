package models

type URLStats struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	ShortURL string `json:"short_url"`
	Visits   int32  `json:"visits"`
}
