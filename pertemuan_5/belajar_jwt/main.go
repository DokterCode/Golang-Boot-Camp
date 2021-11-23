package main

import (
	"belajar_jwt/api"
	"belajar_jwt/models"
	"belajar_jwt/repositories"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Initiasi data
	repositories.Data = []models.UserModel{}
	// Load env file
	godotenv.Load(".env")
	// Server
	ginServer := gin.Default()
	api.RegisterEndpoints(ginServer)
	ginServer.Run()
}
