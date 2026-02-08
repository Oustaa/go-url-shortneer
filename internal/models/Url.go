// Package models implements the app's models
package models

type URL struct {
	ID       int64  `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
