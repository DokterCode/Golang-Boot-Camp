package repositories

import (
	"belajar_jwt/models"
	"errors"
)

// Penyimpan data

var Data []models.UserModel

func AddUser(data models.UserModel) (user models.UserModel, err error) {
	_, err = GetUserByUsername(data.Username)
	if err == nil {
		err = errors.New("Username sudah digunakan")
		return
	}
	err = nil
	Data = append(Data, data)
	user = data
	return
}

func GetUserByUsername(username string) (user *models.UserModel, err error) {
	for _, u := range Data {
		if u.Username == username {
			user = &u
		}
	}
	if user == nil {
		err = errors.New("Data tidak ditemukan")
		return
	}
	return
}

func GetUserByAuth(auth models.AuthModel) (user models.UserModel, err error) {
	return
}
