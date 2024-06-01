package models

type AuthInput struct {
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname"`
	Password string `json:"password" binding:"required"`
}
