package services

import (
	"backendtest-go/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func init() {
	models.LoadEnvs()
	models.ConnectDB()
}

func MarkUnsafeAPI(c *gin.Context) {

	var fileReviewInput models.FileReviewInput
	var user models.User

	err := c.ShouldBindJSON(&fileReviewInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Where("id = ?", fileReviewInput.ReviewerId).Find(&user)
	if !user.IsAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "You need admin user to do this"})
		return
	}

	review := models.FileReview{
		FileID:     fileReviewInput.FileID,
		ReviewerId: fileReviewInput.ReviewerId,
		Unsafe:     fileReviewInput.Unsafe,
	}

	models.DB.Create(&review)

	// rabbitMQ
	fileIDStr := fmt.Sprintf("%v", fileReviewInput.FileID)
	err = models.RabbitMQChannel.Publish(
		"",
		models.UnsafeQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fileIDStr),
		},
	)
	if err != nil {
		log.Println("RabbitMQ Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{"response": review})

}

func DeleteFile(msg []byte) {

	fileStr := string(msg)
	fileID, err := strconv.Atoi(fileStr)
	if err != nil {
		log.Println("Error on converting bytes to int:", err)
	}

	log.Println("Received FileID:", fileID)

	var file models.File
	models.DB.Where("id = ?", fileID).Find(&file)
	file.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_ACCESS_KEY"), ""),
	})
	if err != nil {
		log.Println("error:", err)
	}

	svc := s3.New(sess)

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(file.FilePath),
	})
	if err != nil {
		log.Println("Err to delete object:", err)
	}

	err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(file.FilePath),
	})
	if err != nil {
		log.Println("Error to wait object exclude:", err)
	}

	models.DB.Save(&file)

}
