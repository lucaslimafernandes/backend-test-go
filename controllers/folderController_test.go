package controllers

import (
	"backendtest-go/models"
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

	folder := models.Folder{}

}

func TestListFolders(t *testing.T) {}

func BenchmarkCreateFolder(b *testing.B) {}

func BenchmarkListFolders(b *testing.B) {}

func TestFailCreateFolder(t *testing.T) {}

func TestFailListFolders(t *testing.T) {}
