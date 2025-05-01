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

func GetComments(postId uint, db *gorm.DB) ([]map[string]interface{}, error) {
    var comments []models.Comment
    
    if err := db.Where("post = ?", postId).Find(&comments).Error; err != nil {
        return nil, err
    }

    var userIDs []uint
    for _, c := range comments {
        userIDs = append(userIDs, c.User)
    }

    var users []models.User
    if err := db.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
        return nil, err
    }

    userMap := make(map[uint]models.User)
    for _, u := range users {
        userMap[u.ID] = u
    }

    var result []map[string]interface{}
    for _, c := range comments {
        commentData := map[string]interface{}{
            "id":        c.ID,
            "message":  c.Message,
            "user":     userMap[c.User], 
            "post":      c.Post,
            "createdAt": c.CreatedAt,
            "updatedAt": c.UpdatedAt,
        }
        result = append(result, commentData)
    }

    return result, nil
}