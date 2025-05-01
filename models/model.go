package models

import (
	"time"

	"gorm.io/datatypes"
)
type User struct {
	ID uint `gorm:"primaryKey"`
	Name string
	Phone string
	CreatedAt time.Time
	UpdatedAt time.Time
	AuthOtp string
}

type Post struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	UserID       uint      `json:"userId"`  
	Media        datatypes.JSON `gorm:"type:jsonb" json:"media"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	PostType     string    `json:"postType"`
	ShowLocation bool      `json:"showLocation"`
	Latitude     string    `json:"latitude"`
	Longitude    string    `json:"longitude"`
}

type Comment struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Message string `json:"message"`
	User       uint      `json:"user"`  
	Post       uint      `json:"post"`  
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Bookmark struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	User  uint `gorm:"column:user" json:"user"`
	Post  uint `gorm:"column:post" json:"post"`
}
