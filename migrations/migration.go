package migrations

import (
	"fakye-server/models"

	"gorm.io/gorm"
)


func RunDbMigrations(db *gorm.DB) error {
     err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Bookmark{}, &models.Like{})

	 if err != nil {
		return err
	 }

	 return nil
}
