package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var sessions = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Handle connections: ", err)
	}

	defer conn.Close()

	sessions[conn] = true

	// Read from websocket
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Fatal("Error reading from websocket: ", err)
			delete(sessions, conn)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for i := range broadcast {
		for conn := range sessions {
			err := conn.WriteJSON(i)
			if err != nil {
				fmt.Println(err)
				conn.Close()
				delete(sessions, conn)
			}
		}
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../public/")))
	http.HandleFunc("/ws", handleConnections) // Upgrade and handle websocket connections

	// Listen for incoming messages
	go handleMessages()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
