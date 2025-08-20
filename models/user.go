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

type VerifyCode struct {
	Number string `json:"number"`
	Code   string `json:"code"`
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
