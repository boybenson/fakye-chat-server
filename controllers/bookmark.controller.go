package controllers

import (
	"encoding/json"
	"fakye-server/services"
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