package controllers

import (
	"log"
	"net/http"
	"strconv"
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

func GetMahasiswaByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka",
		})
		return
	}
	mahasiswa, err := repositories.GetMahasiswaById(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
	return
}

func UpdateMahasiswaByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka",
		})
		return
	}

	var mahasiswa *models.Mahasiswa    // buat variabel Mahasiswa
	err = c.ShouldBindJSON(&mahasiswa) // Decode json body ke instance mahasiswa
	if err != nil {
		log.Println("Error membaca data ", err.Error())
		c.Status(http.StatusBadRequest)
		c.Writer.WriteString("Data yang anda kirim salah")
		return
	}

	mahasiswa, err = repositories.UpdateMahasiswaById(idInt, mahasiswa)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, mahasiswa)
	return
}

func DeleteMahasiswaByIdHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID harus berupa angka",
		})
		return
	}
	err = repositories.DeleteMahasiswaById(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil di hapus",
	})
	return
}
