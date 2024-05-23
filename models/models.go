package models

import "time"

type User struct {
	Id           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	IsPremium    bool      `json:"is_premium"`
	LanguageCode string    `json:"language_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AddUser struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	IsPremium    bool   `json:"is_premium"`
	LanguageCode string `json:"language_code"`
}

type AllUsers struct {
	Users []User `json:"users"`
	Count int64  `json:"count"`
}