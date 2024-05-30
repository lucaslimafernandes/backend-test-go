package main

import (
	"backendtest-go/models"
)

func init() {

	models.LoadEnvs()
	models.ConnectDB()

}

func main() {

	models.DB.AutoMigrate(&models.User{}, &models.Folder{})

}
