package models

type User struct {
	Id       int
	Username string
	Number   string
	Email    string
	Password string
}

type UserSignUp struct {
	Username string `json:"username"`
	Number   string `json:"number" binding:"required"`
	Email    string `json:"email"`
}
