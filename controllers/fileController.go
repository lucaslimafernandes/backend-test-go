package controllers

import (
	"archive/zip"
	"backendtest-go/models"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func FileUpload(c *gin.Context) {

	var fileInput models.FileInput
	var uploadBody io.ReadSeeker

	err := c.ShouldBindHeader(&fileInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file is receveid"})
		return
	}

	src, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the file"})
		return
	}

	compress := false
	filename := aws.String(fileInput.Path + "/" + fileInput.File)

	if fileInput.Compress == "true" {

		compress = true
		filename = aws.String(fileInput.Path + "/" + fileInput.File + ".zip")

		// Read the request body into a buffer
		buf := new(bytes.Buffer)
		_, err := io.Copy(buf, c.Request.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file data"})
			return
		}

		// Create a zip file
		archive, err := os.Create(fileInput.File + ".zip")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create zip archive"})
			return
		}

		zipWriter := zip.NewWriter(archive)
		w, err := zipWriter.Create(fileInput.File)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create zip entry"})
			return
		}

		// Copy the buffer content to the zip entry
		_, err = io.Copy(w, buf)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file data to zip"})
			return
		}

		// Reopen the zip file for reading
		fz, err := os.Open(fileInput.File + ".zip")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open zip file"})
			return
		}

		// Use the zip file as the upload body
		uploadBody = fz

		fz.Close()
		zipWriter.Close()
		archive.Close()
		_ = os.Remove(fileInput.File + ".zip")

	} else {
		uploadBody = bytes.NewReader(src)
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

	// FilePath:    fileInput.Path + "/" + fileInput.File,
	file := models.File{
		File:        fileInput.File,
		Folder:      fileInput.Path,
		FilePath:    *filename,
		UserID:      fileInput.UserID,
		UserEmail:   fileInput.UserEmail,
		Description: fileInput.Description,
		Compression: compress,
		Unsafe:      false,
		FileUrl:     os.Getenv("S3_FILEPOINT") + *filename,
	}

	uploadParams := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    filename,
		Body:   uploadBody,
	}

	_, err = svc.PutObject(uploadParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	models.DB.Create(&file)

	message := fmt.Sprintf("New file: %s\nuser:%s\ndescription:%s\npath:%s\nurl:%s\n", file.File, file.UserEmail, file.Description, file.Folder, file.FileUrl)
	err = models.RabbitMQChannel.Publish(
		"",
		models.NotifyQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Println("RabbitMQ Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{"data": file})

}

func ListFiles(c *gin.Context) {

	var folder models.FileList
	err := c.ShouldBindHeader(&folder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	svc := s3.New(sess)
	res, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String("bt"), Prefix: aws.String(folder.Folder + "/")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var files []string
	var _text string
	for _, v := range res.Contents {
		strPtr := v.Key
		s0 := strings.Index(*v.Key, "/")
		if s0 != -1 {
			_text = (*strPtr)[s0+1:]
		} else {
			_text = *strPtr
		}
		if _text != ".emptyFolderPlaceholder" {
			files = append(files, _text)
		}

	}

	c.JSON(http.StatusOK, gin.H{"files": files})

}

func ListFilesV2(c *gin.Context) {

	var sl []models.File
	var folderInput models.FileList

	err := c.ShouldBindJSON(&folderInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Where("folder=? and unsafe=false", folderInput.Folder).Find(&sl)
	c.JSON(http.StatusOK, gin.H{"data": sl})

}

func StreamFile(c *gin.Context) {

	var userInput models.StreamInput
	var streamHistory models.StreamHistory
	var file models.File

	err := c.ShouldBindHeader(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileKey := c.Param("filekey")
	rangeHeader := c.GetHeader("Range")

	models.DB.Where("id=? and unsafe=false", fileKey).Find(&file)

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

	var input *s3.GetObjectInput
	if rangeHeader != "" {
		input = &s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    aws.String(file.FilePath),
			Range:  aws.String(rangeHeader),
		}
	} else {
		input = &s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    aws.String(fileKey),
		}
	}

	res, err := svc.GetObject(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file"})
		return
	}
	defer res.Body.Close()

	// Set headers for streaming
	c.Header("Content-Type", *res.ContentType)
	c.Header("Content-Length", fmt.Sprintf("%d", *res.ContentLength))

	if rangeHeader != "" {
		c.Header("Accept-Ranges", "bytes")
		c.Header("Content-Range", *res.ContentRange)
		c.Status(http.StatusPartialContent)
	} else {
		c.Status(http.StatusOK)
	}

	fid, err := strconv.Atoi(fileKey)
	if err != nil {
		log.Println(err)
	}
	streamHistory.FileID = uint(fid)
	streamHistory.UserID = userInput.UserID
	streamHistory.ViewedAt = time.Now()
	models.DB.Create(&streamHistory)

	// Response body
	_, err = c.Writer.Write([]byte(fmt.Sprintf("Content-Length: %d", *res.ContentLength)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stream file to response"})
		return
	}

}
