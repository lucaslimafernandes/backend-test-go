package controllers

import (
	"backendtest-go/models"
	"net/http"
	"os"

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

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),   // Substitua pela sua região
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")), // Substitua pelo seu endpoint, se necessário
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_KEYID"), os.Getenv("S3_KEY"), ""),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cria um novo cliente do S3
	svc := s3.New(sess)

	uploadParams := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")), // Substitua pelo nome do seu bucket
		Key:    aws.String(folderInput.Folder),     // Substitua pelo nome que você deseja dar à imagem no S3
		Body:   nil,
	}

	folder := models.Folder{
		Folder:    folderInput.Folder,
		UserID:    folderInput.UserID,
		UserEmail: folderInput.UserEmail,
	}
	models.DB.Create(&folder)

	_, err = svc.PutObject(uploadParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// todo persistir os dados

}
