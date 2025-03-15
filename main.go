package main

import (
	"data-parameter/config"
	"data-parameter/models"
	"data-parameter/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env-local")
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.LookupValue{})
	r := routes.SetupRouter()
	r.Run(os.Getenv("APP_PORT"))
}
