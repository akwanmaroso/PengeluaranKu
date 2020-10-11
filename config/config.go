package config

import (
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
)

var (
	// DBDriver is name for driver db
	DBDriver = ""
	// DBURL is url to connect db
	DBURL     = ""
	SecretKey []byte
)

// Load configuration file
func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	DBDriver = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("API_SECRET"))
}
