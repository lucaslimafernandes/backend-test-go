package main

import (
	"backendtest-go/controllers"
	"backendtest-go/middlewares"
	"backendtest-go/models"

	"github.com/gin-gonic/gin"
)

func init() {

	models.LoadEnvs()
	models.ConnectDB()

}

func main() {

	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)

	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.GET("/user/isauth", middlewares.CheckAuth, controllers.IsAuth)

	router.POST("/api/folder", middlewares.CheckAuth, controllers.CreateFolder)

	router.Run()

}
