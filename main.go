package main

import (
	"fakye-server/controllers"
	"fakye-server/database"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	port := ":9000"
 
	db, err := database.ConnectDB()


	if err != nil {
		log.Fatal("Unable to connect to the database", err)
	}

	fmt.Println("App is running on port", port)

	http.HandleFunc("/register", controllers.RegisterHandler(db))
	http.HandleFunc("/signin", controllers.SignInHandler(db))
	http.HandleFunc("/verify-otp", controllers.VerifyOtp(db))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
	
}
