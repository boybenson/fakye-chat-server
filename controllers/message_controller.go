package controllers

import (
	"fakye-server/services"
	"net/http"
)

func SendMessage(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	response := services.SendMessage()

	w.Write([]byte(response))
}