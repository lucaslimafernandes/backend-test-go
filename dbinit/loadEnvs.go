package dbinit

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvs() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

}
