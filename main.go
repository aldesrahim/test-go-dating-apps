package main

import (
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	requiredVariables := []string{
		"APP_HOST",
		"JWT_SECRET",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}

	for _, variable := range requiredVariables {
		if os.Getenv(variable) == "" {
			panic("Missing environment variable: " + variable)
		}
	}
}
