package controllers

import (
	"backendtest-go/models"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func init() {

	models.LoadEnvs()
	models.ConnectDB()
	models.RConn()

}

func TestCreateFolder(t *testing.T) {

	folderInput := models.FolderInput{
		Folder:    "Testing",
		UserID:    16,
		UserEmail: "test@test.com",
	}

	var folderFound models.Folder
	models.DB.Where("folder=?", folderInput.Folder).Find(&folderFound)

	if folderFound.ID != 0 {
		t.Error("Folder already exists!")
	}

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_ACCESS_KEY"), ""),
	})

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

	_, err := svc.PutObject(uploadParams)
	if err != nil {
		t.Error("error to create folder!")
	}

	models.DB.Create(&folder)

}

func TestListFolders(t *testing.T) {

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_ACCESS_KEY"), ""),
	})

	svc := s3.New(sess)

	res, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(os.Getenv("S3_BUCKET"))})
	if err != nil {
		t.Error("ListObjects Error", err)
	}

	if res == nil {
		t.Error("Response is nil")
	}

}

func BenchmarkCreateFolder(b *testing.B) {

	for i := 0; i < b.N; i++ {

		folderInput := models.FolderInput{
			Folder:    fmt.Sprintf("Benchmark%v", i),
			UserID:    16,
			UserEmail: "test@test.com",
		}

		var folderFound models.Folder
		models.DB.Where("folder=?", folderInput.Folder).Find(&folderFound)

		if folderFound.ID != 0 {
			b.Error("Folder already exists!")
		}

		sess, _ := session.NewSession(&aws.Config{
			Region:      aws.String(os.Getenv("S3_REGION")),
			Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_ACCESS_KEY"), ""),
		})

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

		_, err := svc.PutObject(uploadParams)
		if err != nil {
			b.Error("error to create folder!")
		}

		models.DB.Create(&folder)

	}

}

func BenchmarkListFolders(b *testing.B) {

	for i := 0; i < b.N; i++ {

		sess, _ := session.NewSession(&aws.Config{
			Region:      aws.String(os.Getenv("S3_REGION")),
			Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_ACCESS_KEY"), ""),
		})

		svc := s3.New(sess)

		res, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(os.Getenv("S3_BUCKET"))})
		if err != nil {
			b.Error("ListObjects Error", err)
		}

		if res == nil {
			b.Error("Response is nil")
		}

	}

}

func TestFailCreateFolder(t *testing.T) {

	folderInput := models.FolderInput{
		Folder:    "Testing",
		UserID:    16,
		UserEmail: "test@test.com",
	}

	var folderFound models.Folder
	models.DB.Where("folder=?", folderInput.Folder).Find(&folderFound)

	if folderFound.ID == 0 {
		t.Error("Folder doesn't exists!")
	}

}
