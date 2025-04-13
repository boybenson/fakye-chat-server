package controllers

import (
	"encoding/json"
	"fakye-server/services"
	"fakye-server/types"
	"net/http"

	"gorm.io/gorm"
)




func CreatePostHandler (db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var body types.CreatePostRequest
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		
		result, error := services.CreatePost(body, db)
		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}

func GetPostsHandler (db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		result, error := services.GetPosts(db)

		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}