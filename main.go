package main

import (
	"backendtest-go/controllers"
	"backendtest-go/dbinit"
	"backendtest-go/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {

	dbinit.LoadEnvs()
	dbinit.ConnectDB()

}

func main() {

	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	router.Run()

}
