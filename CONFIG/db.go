package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
psqlconn := os.Getenv("DATABASE_URL")
db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
if err != nil {
	log.Fatal("Failed to connect the database", err)
}

DB = db

fmt.Println("Connected to database")

}