package models

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Email     string `json:"email" gorm:"unique"`
	Fullname  string `json:"fullname"`
	Password  string `json:"password"`
	CreatedAt time.Time
}
