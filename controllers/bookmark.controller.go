package controllers

import (
	"encoding/json"
	"fakye-server/services"
	"fakye-server/types"
	"net/http"

	"gorm.io/gorm"
)

func GetBookmarksHandler (db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userId := r.URL.Query().Get("user-id")

		if userId == "" {
			http.Error(w, "Query parameter 'user-id' is required", http.StatusBadRequest)
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

func IsPostBookmarkedHandler(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body types.IsBookMarkRequest
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}


		result, error := services.IsPostBookmarked(body, db)

		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}