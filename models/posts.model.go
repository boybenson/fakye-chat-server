package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	UserID       string    `json:"userId"`
	Media        []string  `json:"media"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	PostType     string    `json:"postType"`
	ShowLocation bool      `json:"showLocation"`
	Latitude     sql.NullString     `json:"latitude"`
	Longitude    sql.NullString    `json:"longitude"`
}
