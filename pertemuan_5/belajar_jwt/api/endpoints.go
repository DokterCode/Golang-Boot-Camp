package api

import (
	"belajar_jwt/auth"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(ginServer *gin.Engine) {
	authMiddleware := auth.AuthMiddleware{}
	// Endpoint untuk registrasi
	ginServer.POST("/register", authMiddleware.RegisterHandler)
	// Endpoint untuk login
	ginServer.POST("/login", authMiddleware.LoginHandler)
	// Endpoint untuk test authentikasi
	ginServer.GET("/testauth")
}
