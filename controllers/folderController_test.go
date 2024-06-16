package controllers

import (
	"backendtest-go/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {

	models.LoadEnvs()
	models.ConnectDB()
	models.RConn()

}

func TestCreateFolder(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/api/v1/folder", CreateFolder)

	folderInput := models.FolderInput{
		Folder:    "Testing",
		UserID:    16,
		UserEmail: "test@test.com",
	}

	jsonValue, _ := json.Marshal(folderInput)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/folder", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}

func TestListFolders(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/api/v1/folder", ListFolders)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/folder", bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}

func BenchmarkCreateFolder(b *testing.B) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/api/v1/folder", CreateFolder)

	for i := 0; i < b.N; i++ {

		folderInput := models.FolderInput{
			Folder:    fmt.Sprintf("Testing%v", i),
			UserID:    16,
			UserEmail: "test@test.com",
		}

		jsonValue, _ := json.Marshal(folderInput)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/folder", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}

	}

}

func BenchmarkListFolders(b *testing.B) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/api/v1/folder", ListFolders)

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/folder", bytes.NewBuffer([]byte("")))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}

	}

}

func TestFailCreateFolder(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/api/v1/folder", CreateFolder)

	folderInput := models.FolderInput{}

	jsonValue, _ := json.Marshal(folderInput)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/folder", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Error("Expected:", http.StatusBadRequest, "Got:", resp.Code)
	}

}

// The test was not created, because it has no input to validate.
// func TestFailListFolders(t *testing.T) {}
