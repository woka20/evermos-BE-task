package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT     string
	DB_NAME  string
	DB_PORT  string
	DB_PASS  string
	DB_UNAME string
	DB_HOST  string
)

func ConfigInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Getting environment variables for config Failed")
	}

	PORT = os.Getenv("PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_PASS = os.Getenv("DB_PASS")
	DB_UNAME = os.Getenv("DB_USERNAME")
	DB_HOST = os.Getenv("DB_HOST")

}
