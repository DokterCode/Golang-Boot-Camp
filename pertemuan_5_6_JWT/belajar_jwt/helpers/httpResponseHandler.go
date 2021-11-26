package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseHandler struct{}

func (this *ResponseHandler) SendHttpBadRequest(message string, ginC *gin.Context) {
	ginC.JSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
	return
}

func (this *ResponseHandler) SendHttpNotFound(message string, ginC *gin.Context) {
	ginC.JSON(http.StatusNotFound, gin.H{
		"message": message,
	})
	return
}

func (this *ResponseHandler) SendHttpServerError(message string, ginC *gin.Context) {
	ginC.JSON(500, gin.H{
		"message": message,
	})
	return
}

func (this *ResponseHandler) SendHttpCreated(data interface{}, ginC *gin.Context) {
	ginC.JSON(http.StatusCreated, data)
	return
}

func (this *ResponseHandler) SendHttpUnAuthorized(message string, ginC *gin.Context) {
	ginC.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
	})
	return
}

func (this *ResponseHandler) SendHttpOk(data interface{}, ginC *gin.Context) {
	ginC.JSON(http.StatusOK, data)
	return
}
