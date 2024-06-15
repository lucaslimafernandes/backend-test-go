package controllers

import (
	"backendtest-go/middlewares"
	"backendtest-go/models"
	"bytes"
	"encoding/json"
	"io"
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

func TestIsAuth(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/user/isauth", IsAuth)

	user := models.AuthInput{
		Email:    "test1@test.com",
		Password: "password",
	}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/user/isauth", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}

func TestGetUserProfile(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/user/profile", GetUserProfile)

	user := models.AuthInput{
		Email:    "test1@test.com",
		Password: "password",
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/user/profile", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Error("Expected:", http.StatusOK, "Got:", resp.Code)
	}

}

func BenchmarkLogin(b *testing.B) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.POST("/auth/login", Login)

	user := models.AuthInput{
		Email:    "test1@test.com",
		Password: "password",
	}

	jsonValue, _ := json.Marshal(user)

	// benchmark
	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}

	}

}

func BenchmarkIsAuth(b *testing.B) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.POST("/user/isauth", IsAuth)

	user := models.AuthInput{
		Email:    "test1@test.com",
		Password: "password",
	}

	jsonValue, _ := json.Marshal(user)

	// benchmark
	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(http.MethodPost, "/user/isauth", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}
	}

}

func BenchmarkGetUserProfile(b *testing.B) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.POST("/user/profile", GetUserProfile)

	user := models.AuthInput{
		Email:    "test1@test.com",
		Password: "password",
	}

	jsonValue, _ := json.Marshal(user)

	// benchmark
	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(http.MethodPost, "/user/profile", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			b.Error("Expected:", http.StatusOK, "Got:", resp.Code)
		}
	}

}

func TestFailCreateUser(t *testing.T) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.POST("/auth/signup", CreateUser)

	user := models.AuthInput{}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Error("Expected:", http.StatusBadRequest, "Got:", resp.Code)
	}

}

func TestFailLogin(t *testing.T) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.POST("/auth/login", Login)

	user := models.AuthInput{}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Error("Expected:", http.StatusBadRequest, "Got:", resp.Code)
	}

}

func TestFailIsAuth(t *testing.T) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.POST("/user/isauth", middlewares.CheckAuth, IsAuth)

	user := models.AuthInput{}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/user/isauth", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Error("Expected:", http.StatusUnauthorized, "Got:", resp.Code)
	}

}

func TestFailGetUserProfile(t *testing.T) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.POST("/user/profile", middlewares.CheckAuth, IsAuth)

	user := models.AuthInput{}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/user/profile", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Error("Expected:", http.StatusUnauthorized, "Got:", resp.Code)
	}
}
