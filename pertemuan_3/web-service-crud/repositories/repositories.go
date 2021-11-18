package repositories

import (
	"log"
	"web-service-crud/models"
)

// Untuk menyimpan data mahasiswa
var DataMahasiswa []models.Mahasiswa

func AddMahasiswa(newData models.Mahasiswa) models.Mahasiswa {
	newData.ID = len(DataMahasiswa) + 1 // Memberikan id untuk mahasiswa
	DataMahasiswa = append(DataMahasiswa, newData)
	log.Println("Data mahasiswa : ", DataMahasiswa)
	return newData
}

func GetAllMahasiswa() []models.Mahasiswa {
	return DataMahasiswa
}
