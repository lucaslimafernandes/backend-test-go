package main

import (
	"backendtest-go/controllers"
	"backendtest-go/middlewares"
	"backendtest-go/migrate"
	"backendtest-go/models"
	"backendtest-go/services"

	"github.com/gin-gonic/gin"
)

func init() {

	models.ConnectDB()
	models.RConn()
	migrate.Migrate()

}

func main() {

	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)

	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.GET("/user/isauth", middlewares.CheckAuth, controllers.IsAuth)

	router.GET("/api/v1/file", middlewares.CheckAuth, controllers.ListFiles)
	router.GET("/api/v2/file", middlewares.CheckAuth, controllers.ListFilesV2)
	router.GET("/api/v1/folder", middlewares.CheckAuth, controllers.ListFolders)
	router.POST("/api/v1/file", middlewares.CheckAuth, controllers.FileUpload)
	router.POST("/api/v1/folder", middlewares.CheckAuth, controllers.CreateFolder)

	router.POST("/api/v1/unsafe", middlewares.CheckAuth, services.MarkUnsafeAPI)

	router.GET("/stream/:filekey", middlewares.CheckAuth, controllers.StreamFile)
	router.GET("/web/stream/", services.ServeHtmlStream)

	router.Run()

}
