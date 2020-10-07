package config

import (
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
)

var (
	// DbDriver is name for driver db
	DbDriver = ""
	// DbURL is url to connect db
	DbURL      = ""
	SECRET_KEY []byte
)

// Load configuration file
func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DbDriver = os.Getenv("DB_DRIVER")
	DbURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	SECRET_KEY = []byte(os.Getenv("API_SECRET"))
}
