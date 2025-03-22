package main

import (
	"fakye-chat-server/controllers"
	"fmt"
	"log"
	"net/http"
)


func main() {
	port := ":8080"
	fmt.Println("App is running on port", port)

	http.HandleFunc("/send-message", controllers.SendMessage)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
