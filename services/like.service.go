package services

import (
	"errors"
	"fakye-server/models"
	"fakye-server/types"
	"fmt"

	"gorm.io/gorm"
)

func ToggleLike(db *gorm.DB, payload types.ToggleLikeRequest) (bool, error) {
	var like models.Like

	err := db.Where(`"user" = ? AND post = ?`, payload.UserID, payload.PostID).First(&like).Error


	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newLike := models.Like{
				User: payload.UserID,
				Post: payload.PostID,
			}
			if err := db.Create(&newLike).Error; err != nil {
				return false, fmt.Errorf("failed to create like: %w", err)
			}
			return true, nil
		}

		return false, fmt.Errorf("failed to check like existence: %w", err)
	}

	if err := db.Delete(&like).Error; err != nil {
		return false, fmt.Errorf("failed to delete like: %w", err)
	}

	return true, nil
}

func IsPostLiked(payload types.IsPostLiked, db *gorm.DB) (bool, error) {
    var like models.Like
    result := db.Where(`"user" = ? AND "post" = ?`, payload.UserID, payload.PostID).First(&like)
    
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return false, nil 
        }
        return false, result.Error 
    }
    
    return true, nil 
}