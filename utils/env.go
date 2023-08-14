package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetToEnv(value string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Error Load Env : " + err.Error())
	}

	result := os.Getenv(value)

	if result == "" {
		log.Fatal(value + " Not Found in env")
	}

	return result
}
