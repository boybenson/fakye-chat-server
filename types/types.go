package types

import "fakye-server/models"

type RegisterRequest struct {
	Phone string `json:"phone"`
	Name    string `json:"name"`
}


type SignInRequest struct{
	Phone string `json:"phone"`
}


type UserWithToken struct {
	models.User
	Token string `json:"token"`
}
