package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Failed to upgrade connection", err)
		return
	}

	clients[conn] = true

	fmt.Println("Client connected")

	defer func() {
		conn.Close()
		delete(clients, conn)
	}()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		fmt.Println("Received message ", string(message))

		//send message to broadcast channel
		broadcast <- message
	}

}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)

			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}

	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	// Start broadcaster goroutine
	go handleMessages()

	fmt.Println("Server running on 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
