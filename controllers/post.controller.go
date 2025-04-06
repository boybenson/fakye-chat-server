package controllers

import (
	"encoding/json"
	"fakye-server/services"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPostsHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := services.GetPosts(db)
		if err != nil {
			http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}

