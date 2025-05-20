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

type CreatePostRequest struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	UserID       uint   `json:"userId"`
	Media        []string `json:"media"`
	ShowLocation bool     `json:"showLocation"`
	Latitude     string   `json:"latitude"`
	Longitude    string   `json:"longitude"`
}

type CreateCommentRequest struct {
	Message  string   `json:"message"`
	User     uint   `json:"userId"`
	Post     uint   `json:"postId"`
}

type GetBookmarksFilter struct {
	UserID uint   `json:"userId"`
}

type ModerationRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type ToggleBookMarkRequest struct {
	UserID uint   `json:"userId"`
	PostID uint 	`json:"postId"`
}

type ToggleLikeRequest struct {
	UserID uint   `json:"userId"`
	PostID uint 	`json:"postId"`
}

type IsBookMarkRequest struct {
	UserID uint   `json:"userId"`
	PostID uint 	`json:"postId"`
}

