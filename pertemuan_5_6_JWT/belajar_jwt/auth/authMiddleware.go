package auth

import (
	"belajar_jwt/helpers"
	"belajar_jwt/models"
	"belajar_jwt/repositories"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
func (this *AuthMiddleware) Authenticate(ginC *gin.Context) (user *models.UserModel, err error) {
	// Ambil token dari header
	token, err := this.getToken(ginC)
	if err != nil {
		log.Println(err)
		return
	}
	// Claim token
	username, err := this.claimToken(token)
	if err != nil {
		log.Println(err)
		return
	}
	// Ambil user dengan username dari token
	user, err = repositories.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (this *AuthMiddleware) getToken(ginC *gin.Context) (token string, err error) {
	token = ginC.Request.Header.Get("Authorization")
	// Validasi bearer token
	token, isBearerToken := this.isBearerToken(token)
	if !isBearerToken {
		err = errors.New("Token tidak valid")
		return
	}
	return
}

func (this *AuthMiddleware) isBearerToken(token string) (tokenWithoutBearer string, isBearer bool) {
	// Memvalidasi token adalah sebuah Bearer token
	if strings.Contains(strings.ToLower(token), "bearer ") {
		isBearer = true
	}
	// Ambil token tanpa bearer jika token adalah sebuah bearer token
	if isBearer {
		tokenWithoutBearer = strings.Split(token, " ")[1]
	}
	return
}

// Validasi token jwt dengan secret_key
func (this *AuthMiddleware) isValidToken(tokenString string) (token *jwt.Token, err error) {
	// Parse token untuk memvalidasi, jika valid akan mengembalikan secret key yang sesuai
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return
	}
	return
}

// Ambil data username dari token
func (this *AuthMiddleware) claimToken(tokenString string) (username string, err error) {
	token, err := this.isValidToken(tokenString)
	if err != nil {
		// Kebutuhan debuging
		log.Println(err)
		// Client perpose
		err = errors.New("Token tidak valid")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	// Jika claim berhasil
	if ok && token.Valid {
		log.Println("User auth :=>", claims["user_id"])
		username = fmt.Sprintf("%v", claims["user_id"])
	}
	return
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
