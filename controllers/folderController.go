package controllers

import (
	"backendtest-go/models"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func CreateFolder(c *gin.Context) {

	var folderInput models.FolderInput

	err := c.ShouldBindJSON(&folderInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var folderFound models.Folder
	models.DB.Where("folder=?", folderInput.Folder).Find(&folderFound)

	if folderFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this folder already exists"})
		return
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_ACCESS_KEY"), ""),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cria um novo cliente do S3
	svc := s3.New(sess)

	folder := models.Folder{
		Folder:    folderInput.Folder,
		UserID:    folderInput.UserID,
		UserEmail: folderInput.UserEmail,
	}

	uploadParams := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(folderInput.Folder + "/"),
		Body:   nil,
	}

	// Faz o upload da imagem
	_, err = svc.PutObject(uploadParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&folder)

	c.JSON(http.StatusOK, gin.H{"ok": folder})

}

func ListFolders(c *gin.Context) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_ACCESS_KEY"), ""),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	svc := s3.New(sess)
	res, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String("bt")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var folders []string
	var _text string
	for _, v := range res.Contents {
		strPtr := v.Key
		s0 := strings.Index(*v.Key, "/")
		if s0 != -1 {
			_text = (*strPtr)[:s0]
		} else {
			_text = *strPtr
		}
		if !slices.Contains(folders, _text) {
			folders = append(folders, _text)
		}

	}

	c.JSON(http.StatusOK, gin.H{"folders": folders})

}
