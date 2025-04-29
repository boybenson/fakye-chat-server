package services

import (
	"errors"
	"fakye-server/models"
	"fakye-server/types"
	"fmt"

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



func ToggleBookMark(db *gorm.DB, payload types.ToggleBookMarkRequest) (bool, error) {
    var bookmark models.Bookmark

    err := db.Where("user = ? AND post = ?", "15", "10").First(&bookmark).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            newBookmark := models.Bookmark{
                User: payload.UserID,
                Post: payload.PostID,
            }
            if err := db.Create(&newBookmark).Error; err != nil {
                return false, fmt.Errorf("failed to create bookmark: %w", err)
            }
            return true, nil 
        }
        return false, fmt.Errorf("failed to check bookmark existence: %w", err)
    }

    if err := db.Delete(&bookmark).Error; err != nil {
        return false, fmt.Errorf("failed to delete bookmark: %w", err)
    }

    return false, nil 
}