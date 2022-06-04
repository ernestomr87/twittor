package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvs(variable string) string {
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)

	}
	return envs[variable]
}
