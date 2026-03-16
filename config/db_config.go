package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Check if the file exists in the root folder.")
	}

	database, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = database
	return nil
}
