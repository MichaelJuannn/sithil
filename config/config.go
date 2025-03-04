package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error env loading from .env file trying to load from OS")
	}
	return os.Getenv(key)
}
