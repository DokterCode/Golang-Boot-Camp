package controllers

import (
	"log"
	"net/http"
	"web-service-crud/models"
	"web-service-crud/repositories"

	"github.com/gin-gonic/gin"
)

func AddMahasiswaHandler(c *gin.Context) {

	// Decode json body ke instance dari model Mahasiswa
	var newMahasiswa models.Mahasiswa      // buat variabel Mahasiswa
	err := c.ShouldBindJSON(&newMahasiswa) // Decode json body ke instance mahasiswa
	if err != nil {
		log.Println("Error membaca data ", err.Error())
		c.Status(http.StatusBadRequest)
		c.Writer.WriteString("Data yang anda kirim salah")
		return
	}
	// log.Println("Data mahasiswa : ", newMahasiswa)
	newMahasiswa = repositories.AddMahasiswa(newMahasiswa)
	// Berikan status 'created : 201' untuk data yang berhasil dibuat
	c.JSON(http.StatusCreated, newMahasiswa)
}

func GetAllMahasiswaHandler(c *gin.Context) {
	data := repositories.GetAllMahasiswa()
	c.JSON(http.StatusOK, data)
}
