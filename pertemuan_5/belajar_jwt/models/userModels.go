package models

type UserModel struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"` // not required
	Password string `json:"password" binding:"required`
}
