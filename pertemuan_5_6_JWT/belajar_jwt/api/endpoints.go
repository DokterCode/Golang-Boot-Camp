package api

import (
	"belajar_jwt/auth"
	"belajar_jwt/helpers"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(ginServer *gin.Engine) {
	authMiddleware := auth.AuthMiddleware{}
	// Endpoint untuk registrasi
	ginServer.POST("/register", authMiddleware.RegisterHandler)
	// Endpoint untuk login
	ginServer.POST("/login", authMiddleware.LoginHandler)
	// Endpoint untuk test authentikasi
	ginServer.GET("/testauth", func(c *gin.Context) {
		responseHandler := helpers.ResponseHandler{}
		authMiddleware := auth.AuthMiddleware{}
		_, err := authMiddleware.Authenticate(c)
		// Jika tidak terauthentikasi
		if err != nil {
			responseHandler.SendHttpUnAuthorized(err.Error(), c)
			return
		}
		responseHandler.SendHttpOk("Berhasil authentikasi dengan token", c)
		return
	})
}
