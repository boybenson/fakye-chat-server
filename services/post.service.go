package services

import (
	"encoding/json"
	"fakye-server/models"
	"fakye-server/types"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)




func CreatePost(payload types.CreatePostRequest, db *gorm.DB) (bool, error) {

	mediaJSON, err := json.Marshal(payload.Media)
	if err != nil {
		return false, err
	}

	post := models.Post{
		Latitude:      payload.Latitude,
		Longitude:     payload.Longitude,
		Name:          payload.Name,
		Description:   payload.Description,
		PostType:      payload.Type,
		UserID:        payload.UserID,
		ShowLocation:  payload.ShowLocation,
		Media:         datatypes.JSON(mediaJSON),
	}
	result := db.Create(&post)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func GetPosts(db *gorm.DB)([]models.Post, error){
	var posts []models.Post
	result := db.Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}