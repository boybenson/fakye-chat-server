package migrations

import (
	"fakye-server/models"

	"gorm.io/gorm"
)


func RunDbMigrations(db *gorm.DB) error {
     err := db.AutoMigrate(&models.User{}, &models.Post{})

	 if err != nil {
		return err
	 }

	 return nil
}
