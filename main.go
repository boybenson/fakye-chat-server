package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetUsers (w http.ResponseWriter, r *http.Request){
}

func main(){
	port := ":8080"
	fmt.Println("App is running on port", port)
	err := http.ListenAndServe(port, nil)

	http.HandleFunc("/app", GetUsers)

	if err != nil {
		log.Fatal(err)
	}

}