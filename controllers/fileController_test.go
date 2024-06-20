package controllers

import (
	"backendtest-go/models"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {

	models.LoadEnvs()
	models.ConnectDB()
	models.RConn()

}

func TestFileUpload(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/api/v1/file", FileUpload)

	fileInput := models.FileInput{
		File:        "Testing.mp4",
		Path:        "Testing",
		UserID:      3,
		UserEmail:   "test@test.com",
		Description: "Testing",
	}

	file, _ := os.Open("15MB.mp4")
	defer file.Close()

	fileContent, _ := io.ReadAll(file)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/file", bytes.NewBuffer(fileContent))
	req.Header.Set("file", fileInput.File)
	req.Header.Set("path", fileInput.Path)
	req.Header.Set("userid", fmt.Sprintf("%v", fileInput.UserID))
	req.Header.Set("useremail", fileInput.UserEmail)
	req.Header.Set("description", fileInput.Description)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}

func TestListFiles(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/api/v1/file", ListFiles)

	folder := models.FileList{
		Folder: "Testing",
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/file", bytes.NewBuffer([]byte("")))
	req.Header.Set("folder", folder.Folder)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}

func TestListFilesV2(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/api/v2/file", ListFiles)

	folder := models.FileList{
		Folder: "Testing",
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/v2/file", bytes.NewBuffer([]byte("")))
	req.Header.Set("folder", folder.Folder)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}

// Not implemented
// func TestStreamFile(t *testing.T) {

// 	gin.SetMode(gin.TestMode)

// 	router := gin.Default()
// 	router.GET("/stream/:filekey", StreamFile)

// 	req, _ := http.NewRequest(http.MethodGet, "/stream/12", bytes.NewBuffer([]byte("")))
// 	req.Header.Set("userid", "16")
// 	req.Header.Set("Range", "bytes=0-1023")

// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	if resp.Code != http.StatusPartialContent {
// 		t.Error("Expected:", http.StatusPartialContent, "Got:", resp.Code)
// 	}

// }

func BenchmarkFileUpload(b *testing.B) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/api/v1/file", FileUpload)

	for i := 0; i < b.N; i++ {

		fileInput := models.FileInput{
			File:        fmt.Sprintf("Testing%v.mp4", i),
			Path:        "Benchmark",
			UserID:      3,
			UserEmail:   "test@test.com",
			Description: "Testing",
		}

		file, _ := os.Open("15MB.mp4")
		defer file.Close()

		fileContent, _ := io.ReadAll(file)

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/file", bytes.NewBuffer(fileContent))
		req.Header.Set("file", fileInput.File)
		req.Header.Set("path", fileInput.Path)
		req.Header.Set("userid", fmt.Sprintf("%v", fileInput.UserID))
		req.Header.Set("useremail", fileInput.UserEmail)
		req.Header.Set("description", fileInput.Description)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}
	}

}

func BenchmarkListFiles(b *testing.B) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/api/v1/file", ListFiles)

	folder := models.FileList{
		Folder: "Testing",
	}

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/file", bytes.NewBuffer([]byte("")))
		req.Header.Set("folder", folder.Folder)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}
	}

}

func BenchmarkListFilesV2(b *testing.B) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/api/v2/file", ListFiles)

	folder := models.FileList{
		Folder: "Testing",
	}

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, "/api/v2/file", bytes.NewBuffer([]byte("")))
		req.Header.Set("folder", folder.Folder)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}

	}

}

func BenchmarkStreamFile(b *testing.B) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/stream/:filekey", StreamFile)

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, "/stream/12", bytes.NewBuffer([]byte("")))
		req.Header.Set("userid", "16")
		req.Header.Set("Range", "bytes=0-1023")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusPartialContent {
			b.Error("Expected:", http.StatusPartialContent, "Got:", resp.Code)
		}

	}

}

func TestFailFileUpload(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/api/v1/file", FileUpload)

	fileInput := models.FileInput{
		File:        "Testing.mp4",
		Path:        "Testing",
		UserID:      3,
		UserEmail:   "test@test.com",
		Description: "Testing",
	}

	file, _ := os.Open("15MB.mp4")
	defer file.Close()

	fileContent, _ := io.ReadAll(file)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/file", bytes.NewBuffer(fileContent))
	req.Header.Set("file", fileInput.File)
	req.Header.Set("path", fileInput.Path)
	req.Header.Set("description", fileInput.Description)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Error("Expected:", http.StatusBadRequest, "Got:", resp.Code)
	}

}

// The test was not created, because it has no input to validate.
// func TestFailListFiles(t *testing.T) {}

// The test was not created, because it has no input to validate.
// func TestFailListFilesV2(t *testing.T) {}

func TestFailStreamFile(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/stream/:filekey", StreamFile)

	req, _ := http.NewRequest(http.MethodGet, "/stream/12", bytes.NewBuffer([]byte("")))
	req.Header.Set("Range", "bytes=0-1023")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Error("Expected:", http.StatusBadRequest, "Got:", resp.Code)
	}

}
