package main

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/models"
	"aldesrahim/test-go-dating-apps/routes"
	"github.com/gin-gonic/gin"
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

	database.Connect()
	database.DB.AutoMigrate(
		&models.Product{},
		&models.ProductPeriod{},
		&models.User{},
		&models.Swipe{},
		&models.SwipeHistory{},
		&models.Match{},
		&models.Payment{},
		&models.Subscribe{},
	)

	r := gin.Default()
	routes.ApiRoutes(r)

	r.Run(os.Getenv("APP_HOST"))
}
