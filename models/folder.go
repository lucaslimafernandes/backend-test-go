package models

import "gorm.io/gorm"

type FolderInput struct {
	Folder    string `json:"folder" binding:"required"`
	UserID    int    `json:"userid" binding:"required"`
	UserEmail string `json:"useremail" binding:"required"`
}

type Folder struct {
	gorm.Model
	Folder    string `json:"folder" gorm:"unique"`
	UserID    int    `json:"userid"`
	UserEmail string `json:"useremail"`
}
