package controllers

import (
	"encoding/json"
	"fakye-server/services"
	"fakye-server/types"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func CreateCommentHandler (db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var body types.CreateCommentRequest
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		
		result, error := services.CreateComment(body, db)
		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}

func GetCommentsHandler (db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		postIdStr := r.URL.Query().Get("postId")

		if postIdStr == "" {
			http.Error(w, "Query parameter 'postId' is required", http.StatusBadRequest)
			return
		}

		postIdUint64, err := strconv.ParseUint(postIdStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid postId", http.StatusBadRequest)
			return
		}

		postId := uint(postIdUint64)

		result, error := services.GetComments(uint(postId), db)

		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}