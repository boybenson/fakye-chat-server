package services

import (
	"context"
	"fakye-server/models"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func updateName(x *string){
	*x = "Yeboah"
}

func GetPosts (db *pgxpool.Pool) ([]models.Post, error){

	name := "Benson"

	updateName(&name)

	println(name)


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

func CreatePost (db *pgxpool.Pool) bool{
	return true
}