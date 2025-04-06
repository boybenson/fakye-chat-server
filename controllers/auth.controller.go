package controllers

import (
	"encoding/json"
	"fakye-server/services"
	"fakye-server/types"
	"net/http"

	"gorm.io/gorm"
)




func RegisterHandler(db *gorm.DB) http.HandlerFunc {	
	return func(w http.ResponseWriter, r *http.Request) {
		

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body types.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if body.Name == "" || body.Phone == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		result, error := services.Register(body, db)


		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}

func SignInHandler(db *gorm.DB) http.HandlerFunc {	
	return func(w http.ResponseWriter, r *http.Request) {
		

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body types.SignInRequest
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if  body.Phone == "" {
			http.Error(w, "Phone number is required", http.StatusBadRequest)
			return
		}

		result, error := services.SignIn(body, db)


		if(error != nil){
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result) 

	}
}


