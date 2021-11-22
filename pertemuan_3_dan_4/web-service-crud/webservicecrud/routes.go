package webservicecrud

import (
	"web-service-crud/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginServer *gin.Engine) {
	ginServer.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("My Web Service Crud")
	})
	// Endpoint untuk menambah mahasiswa
	ginServer.POST("/mahasiswa", controllers.AddMahasiswaHandler)
	// Endpoint untuk mengambil seluruh data mahasiswa
	ginServer.GET("/mahasiswa", controllers.GetAllMahasiswaHandler)
	// Endpoint untuk mengambil 1 data
	ginServer.GET("/mahasiswa/:id", controllers.GetMahasiswaByIDHandler)
	// Endpoint untuk mengubah 1 data
	ginServer.PUT("/mahasiswa/:id", controllers.UpdateMahasiswaByIDHandler)
	// Endpoint untuk memghapus 1 data
	ginServer.DELETE("/mahasiswa/:id", controllers.DeleteMahasiswaByIdHandler)

}
