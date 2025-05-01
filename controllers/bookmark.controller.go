package controllers

import (
	"encoding/json"
	"fakye-server/services"
	"fakye-server/types"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetBookmarksHandler (db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userId := r.URL.Query().Get("userId")

		if userId == "" {
			http.Error(w, "Query parameter 'userId' is required", http.StatusBadRequest)
			return
		}

		result, error := services.GetBookmarks(userId, db)

		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}

func ToggleBookMark(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body types.ToggleBookMarkRequest
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}


		result, error := services.ToggleBookMark(db, body)

		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}


func IsPostBookmarkedHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userIdStr := r.URL.Query().Get("userId")
		postIdStr := r.URL.Query().Get("postId")

		if userIdStr == "" || postIdStr == "" {
			http.Error(w, "Both userId and postId query parameters are required", http.StatusBadRequest)
			return
		}

		userId, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			http.Error(w, "userId must be a valid number", http.StatusBadRequest)
			return
		}

		postId, err := strconv.ParseUint(postIdStr, 10, 64)
		if err != nil {
			http.Error(w, "postId must be a valid number", http.StatusBadRequest)
			return
		}

		isBookmarked, err := services.IsPostBookmarked(types.IsBookMarkRequest{
			UserID: uint(userId),
			PostID: uint(postId),
		}, db)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Not found", http.StatusNotFound)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		json.NewEncoder(w).Encode(map[string]bool{
			"isBookmarked": isBookmarked,
		})
	}
}