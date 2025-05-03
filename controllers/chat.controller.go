package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Upgrading connection error:", err)
        http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
        return
    }
	
    defer conn.Close()

    fmt.Println("Client connected to WebSocket")

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                fmt.Printf("WebSocket error: %v\n", err)
            }
            break
        }

        fmt.Printf("Received message: %s\n", string(message))

        time.Sleep(3 * time.Second)

        if err := conn.WriteMessage(messageType, message); err != nil {
            fmt.Println("Error writing message:", err)
            break
        }
    }
}