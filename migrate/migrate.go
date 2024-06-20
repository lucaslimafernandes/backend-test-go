package migrate

import (
	"backendtest-go/models"
)

func init() {

	// models.LoadEnvs()
	models.ConnectDB()

}

func Migrate() {

	models.DB.AutoMigrate(
		&models.User{},
		&models.Folder{},
		&models.File{},
		&models.FileReview{},
		&models.StreamHistory{},
	)

}
