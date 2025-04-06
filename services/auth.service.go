package services

import (
	"fakye-server/models"
	"fakye-server/types"

	"gorm.io/gorm"
)




func Register(payload types.RegisterRequest, db *gorm.DB)(*models.User, error) {
	user := models.User{Name: payload.Name, Phone: payload.Phone}

	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}