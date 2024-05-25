package main

import (
	"backendtest-go/dbinit"
	"backendtest-go/models"
)

func init() {

	dbinit.LoadEnvs()
	dbinit.ConnectDB()

}

func main() {

	dbinit.DB.AutoMigrate(&models.User{})

}
