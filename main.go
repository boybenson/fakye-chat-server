package main

import (
	"context"
	"fakye-server/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	port := ":9000"
 
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))


	if err != nil {
		log.Fatal("Unable to connect to the database", err)
	}

	defer db.Close()

	fmt.Println("App is running on port", port)

	http.HandleFunc("/get-posts", controllers.GetPostsHandler(db))


	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
	
}
