package services

import (
	"fakye-server/models"

	"gorm.io/gorm"
)

func GetBookmarks (userId string, db *gorm.DB)([]models.Post, error){

	var posts []models.Post
	result := db.Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

