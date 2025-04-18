package services

import (
	"fakye-server/models"
	"fakye-server/types"

	"gorm.io/gorm"
)

func CreateComment(payload types.CreateCommentRequest, db *gorm.DB) (bool, error) {

	comment := models.Comment{
		Message:  payload.Message,
		User: payload.User,
		Post: payload.Post,
	}

	result := db.Create(&comment)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil

}