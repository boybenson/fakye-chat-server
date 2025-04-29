package routes

import (
	"fakye-server/controllers"
	"fakye-server/database"
	"log"
	"net/http"
)

func RootRouter ()(*http.ServeMux, error){
	db, err := database.ConnectDB()
	mux := http.NewServeMux()

	

	if err != nil {
		log.Fatal("Unable to connect to the database", err)
	}

	mux.HandleFunc("/register", controllers.RegisterHandler(db))
	mux.HandleFunc("/signin", controllers.SignInHandler(db))
	mux.HandleFunc("/verify-otp", controllers.VerifyOtp(db))
	mux.HandleFunc("/create-post", controllers.CreatePostHandler(db))
	mux.HandleFunc("/get-posts", controllers.GetPostsHandler(db))



	// Bookmarks
	mux.HandleFunc("/get-bookmarks", controllers.GetBookmarksHandler(db))
	mux.HandleFunc("/toggle-bookmark", controllers.ToggleBookMark(db))

	return mux, nil
}