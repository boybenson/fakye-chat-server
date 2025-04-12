package types

import "fakye-server/models"

type RegisterRequest struct {
	Phone string `json:"phone"`
	Name    string `json:"name"`
}


type SignInRequest struct{
	Phone string `json:"phone"`
}

type VerifyOtpRequest struct{
	Phone string `json:"phone"`
	Otp string `json:"otp"`
}

type UserWithToken struct {
	models.User
	Token string `json:"token"`
}


type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
