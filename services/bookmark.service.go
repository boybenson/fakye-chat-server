package services

import (
	"fakye-server/models"
	"fakye-server/types"

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
    bookmark := models.Bookmark{User: payload.UserID, Post: payload.PostID}

    res := db.First(&bookmark)
    
    if res.Error != nil {
        println("Error:", res.Error.Error())
    } else {
        println("Found bookmark")
    }
    
    return true, nil
}