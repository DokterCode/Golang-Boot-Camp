package main

import (
	"web-service-crud/models"
	"web-service-crud/repositories"
	"web-service-crud/webservicecrud"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi server gin
	ginServer := gin.Default()
	// Inisialisasi data mahasisa kosong
	repositories.DataMahasiswa = []*models.Mahasiswa{}
	webservicecrud.RegisterRoutes(ginServer)
	ginServer.Run()
}
