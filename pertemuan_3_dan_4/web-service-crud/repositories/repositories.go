package repositories

import (
	"errors"
	"log"
	"web-service-crud/models"
)

// Untuk menyimpan data mahasiswa
var DataMahasiswa []*models.Mahasiswa

func AddMahasiswa(newData models.Mahasiswa) models.Mahasiswa {
	newData.ID = len(DataMahasiswa) + 1 // Memberikan id untuk mahasiswa
	DataMahasiswa = append(DataMahasiswa, &newData)
	log.Println("Data mahasiswa : ", DataMahasiswa)
	return newData
}

func GetAllMahasiswa() []*models.Mahasiswa {
	return DataMahasiswa
}

func GetMahasiswaById(id int) (mahasiswa *models.Mahasiswa, err error) {
	// Mencari data mahasiswa berdasarkan ID
	for _, mhs := range DataMahasiswa {
		if mhs.ID == id {
			mahasiswa = mhs
		}
	}
	// Jika tidak ditemukan
	if mahasiswa == nil {
		err = errors.New("Data tidak ditemukan")
	}
	return
}

func UpdateMahasiswaById(id int, data *models.Mahasiswa) (mahasiswa *models.Mahasiswa, err error) {
	mhs, err := GetMahasiswaById(id)
	if err != nil {
		return
	}
	if len(data.Name) > 0 {
		mhs.Name = data.Name
	}
	if len(data.NIM) > 0 {
		mhs.NIM = data.NIM
	}
	for i, m := range DataMahasiswa {
		if m.ID == id {
			DataMahasiswa[i] = mhs
		}
	}
	mahasiswa = mhs
	return
}

func DeleteMahasiswaById(id int) (err error) {
	_, err = GetMahasiswaById(id)
	if err != nil {
		return
	}
	var indexToDelete int
	for i, mhs := range DataMahasiswa {
		if mhs.ID == id {
			indexToDelete = i
		}
	}
	// Hapus data
	DataMahasiswa = append(DataMahasiswa[:indexToDelete], DataMahasiswa[indexToDelete+1:]...)
	return
}
