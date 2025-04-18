package main

import (
	"fakye-server/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	port := ":9000"



	mux, _ := routes.RootRouter()

	

	fmt.Println("App is running on port", port)

	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
