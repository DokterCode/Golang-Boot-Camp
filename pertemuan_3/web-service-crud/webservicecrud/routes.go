package webservicecrud

import (
	"web-service-crud/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginServer *gin.Engine) {
	ginServer.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("My Web Service Crud")
	})
	// Path untuk menambah mahasiswa
	ginServer.POST("/mahasiswa", controllers.AddMahasiswaHandler)
	// Path untuk mengambil seluruh data mahasiswa
	ginServer.GET("/mahasiswa", controllers.GetAllMahasiswaHandler)
}
