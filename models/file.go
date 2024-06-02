package models

import "gorm.io/gorm"

type FileInput struct {
	File        string `json:"file" binding:"required"`
	Path        string `json:"path" binding:"required"`
	UserID      int    `json:"userid" binding:"required"`
	UserEmail   string `json:"useremail" binding:"required"`
	Description string `json:"description"`
	Compress    string `json:"compress"`
}

type File struct {
	gorm.Model
	File        string `json:"file"`
	Folder      string `json:"folder"`
	FilePath    string `json:"filepath" gorm:"unique"`
	UserID      int    `json:"userid"`
	UserEmail   string `json:"useremail"`
	Description string `json:"description"`
	Compression bool   `json:"compression" gorm:"default=false"`
	Unsafe      bool   `json:"unsafe" gorm:"default=false"`
	FileUrl     string `json:"fileurl"`
}

type FileReview struct {
	gorm.Model
	FileID     uint `json:"fileid"`
	ReviewerId uint `json:"reviewerid"`
	Unsafe     bool `json:"unsafe" gorm:"default=false"`
}

type FileReviewInput struct {
	FileID     uint `json:"fileid" binding:"required"`
	ReviewerId uint `json:"reviewerid" binding:"required"`
	Unsafe     bool `json:"unsafe"  binding:"required"`
}

type FileList struct {
	Folder string `json:"folder"`
}
