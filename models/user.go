package models

type User struct {
	Id       int
	Username string
	Number   string
	Email    string
	Password string
}

type UserRegister struct {
	Username string `json:"username"`
	Number   string `json:"number" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}
