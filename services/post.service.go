package services

import (
	"context"
	"fakye-server/models"
	"fakye-server/types"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"
)



func GetPosts (db *pgxpool.Pool) ([]models.Post, error){


	rows, err := db.Query(context.Background(), `SELECT * FROM post`)


	if err != nil {
		log.Println("Error fetching posts:", err)
		return nil, fmt.Errorf("could not fetch posts: %w", err)
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID, &post.Name, &post.Description, &post.UserID, &post.Media,
			 &post.CreatedAt, &post.UpdatedAt, &post.PostType,
			&post.ShowLocation, &post.Latitude, &post.Longitude,
		); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func CreatePost (payload types.CreatePostRequest, db *gorm.DB) (bool, error){
	post := models.Post{Latitude: payload.Latitude, Longitude:payload.Longitude,  Media: payload.Media, Name: payload.Name, Description: payload.Description, PostType: payload.Type, UserID: payload.UserID, ShowLocation: payload.ShowLocation}
	result := db.Create(&post)

	if result.Error != nil {
		return false, result.Error
	}
	
	return true, nil
}