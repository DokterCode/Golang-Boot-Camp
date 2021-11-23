package auth

import (
	"belajar_jwt/helpers"
	"belajar_jwt/models"
	"belajar_jwt/repositories"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthMiddleware struct{}

// Fungsi http handler untuk registrasi
func (this *AuthMiddleware) RegisterHandler(ginC *gin.Context) {
	// Error handler
	responseHandler := helpers.ResponseHandler{}
	// Ambil body registrasi
	var newUser models.UserModel
	err := ginC.ShouldBindJSON(&newUser)
	// Handler jika data yang dikirimkan salah
	if err != nil {
		responseHandler.SendHttpBadRequest("Data yang anda masukkan salah", ginC)
		return
	}

	if len(newUser.Name) < 4 {
		responseHandler.SendHttpBadRequest("Nama minimal 4 karakter", ginC)
		return
	}
	if len(newUser.Username) < 6 {
		responseHandler.SendHttpBadRequest("Username  minimal 6 karakter", ginC)
		return
	}
	if len(newUser.Password) < 6 {
		responseHandler.SendHttpBadRequest("Password  minimal 6 karakter", ginC)
		return
	}

	// Buat user baru
	// Hash password untuk user baru
	newUser.Password, err = this.passwordHash(newUser.Password)
	if err != nil {
		responseHandler.SendHttpServerError("Terjadi kesalahan pada server", ginC)
		return
	}
	_, err = repositories.AddUser(newUser)
	if err != nil {
		responseHandler.SendHttpBadRequest(err.Error(), ginC)
		return
	}
	responseHandler.SendHttpCreated(gin.H{"message": "Berhasil registrasi"}, ginC)
	return
}

// Fungsi http handler untuk login
func (this *AuthMiddleware) LoginHandler(ginC *gin.Context) {
	// Error handler
	responseHandler := helpers.ResponseHandler{}
	// Ambil body registrasi
	var authData models.AuthModel
	err := ginC.ShouldBindJSON(&authData)
	// Handler jika data yang dikirimkan salah
	if err != nil {
		responseHandler.SendHttpBadRequest("Data yang anda masukkan salah", ginC)
		return
	}
	// Get user by input username
	user, err := repositories.GetUserByUsername(authData.Username)
	if err != nil {
		responseHandler.SendHttpServerError("Username tidak dikenal", ginC)
		return
	}
	// Validate the password
	err = this.verifyPassword(authData.Password, user.Password)
	if err != nil {
		responseHandler.SendHttpServerError("Password salah", ginC)
		return
	}
	// Generate jwt
	token, err := this.generateToken(user.Username)
	if err != nil {
		log.Println("Error generate token : ", err)
		responseHandler.SendHttpServerError("Terjadi masalah dengan server", ginC)
		return
	}
	responseHandler.SendHttpOk(gin.H{
		"token": token,
	}, ginC)
	return
}

// Helper authentikasi
func (this *AuthMiddleware) Authenticate(ginC *gin.Context) {

}

func (this *AuthMiddleware) passwordHash(password string) (hashedPassword string, err error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return
	}
	hashedPassword = string(hashedPasswordByte)
	return
}

func (this *AuthMiddleware) verifyPassword(password, hashedPassword string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return
}

func (this *AuthMiddleware) generateToken(username string) (string, error) {
	log.Println(os.Environ())
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}
