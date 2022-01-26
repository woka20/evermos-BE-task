package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func ConfigInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Getting environment variables for config Failed")
	}

	PORT = os.Getenv("PORT")

}
