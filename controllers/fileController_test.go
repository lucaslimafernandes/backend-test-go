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
		UserID:      16,
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

func TestListFiles(t *testing.T) {}

func TestListFilesV2(t *testing.T) {}

func TestStreamFile(t *testing.T) {}

func BenchmarkFileUpload(b *testing.B) {}

func BenchmarkListFiles(b *testing.B) {}

func BenchmarkListFilesV2(b *testing.B) {}

func BenchmarkStreamFile(b *testing.B) {}

func TestFailFileUpload(t *testing.T) {}

func TestFailListFiles(t *testing.T) {}

func TestFailListFilesV2(t *testing.T) {}

func TestFailStreamFile(t *testing.T) {}
