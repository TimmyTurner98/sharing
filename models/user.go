package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username"`
	Number   string `json:"number" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}
