package models

type User struct {
	ID           int64  `json:"id" grom:"primarykey"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
