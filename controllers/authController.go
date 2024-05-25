package controllers

import (
	"backendtest-go/dbinit"
	"backendtest-go/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {

	var authInput models.AuthInput

	err := c.ShouldBindJSON(&authInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	dbinit.DB.Where("email=?", authInput.Email).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    authInput.Email,
		Fullname: authInput.Fullname,
		Password: string(passwordHash),
	}

	dbinit.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})

}

func Login(c *gin.Context) {

	var authInput models.AuthInput

	err := c.ShouldBindJSON(&authInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	dbinit.DB.Where("email=?", authInput.Email).Find(&userFound)
	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func GetUserProfile(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(http.StatusOK, gin.H{"user": user})

}
