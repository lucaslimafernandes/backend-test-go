package controllers

import (
	"backendtest-go/models"
	"bytes"
	"encoding/json"
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

func TestCreateUser(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/auth/signup", CreateUser)

	user := models.User{
		Email:    "test1@test.com",
		Password: "password",
		Fullname: "Test1 User",
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Error("Expected:", http.StatusCreated, "Got:", resp.Code)
	}

	// models.DB.Delete(&user)
}

// ok  	backendtest-go/controllers	3.006s

func TestLogin(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/auth/login", Login)

	userInput := models.AuthInput{
		Email:    "test1@test.com",
		Password: "password",
	}

	jsonValue, _ := json.Marshal(userInput)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}
