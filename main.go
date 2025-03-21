package main

import (
	"data-parameter/config"
	"log/slog"

	//"data-parameter/models"
	"data-parameter/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env-local")

	logger := slog.New(config.LogHandler{})
	slog.SetDefault(logger)

	config.ConnectDatabase()
	config.ConnectRedis()
	//config.DB.AutoMigrate(&models.LookupValue{})
	r := routes.SetupRouter()
	r.Run(os.Getenv("APP_PORT"))
}
