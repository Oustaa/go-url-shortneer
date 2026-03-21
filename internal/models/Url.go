// Package models implements the app's models
package models

import "gorm.io/gorm"

type URL struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	ShortURL string `json:"short_url" gorm:"index"`
	LongURL  string `json:"long_url"`
	Visits   int64  `json:"visits"`
	Deleted  gorm.DeletedAt

	UserID *int64 `json:"-" gorm:"index"`
	User   *User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
