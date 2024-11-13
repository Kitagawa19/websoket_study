package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func handleMultiConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	defer delete(clients,ws)
	clients[ws] = true
	err = ws.WriteMessage(websocket.TextMessage, []byte(" Welcome to WebSocket server!"))
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, ws)
			break
		}
		broadcast <- string(msg)
	}
}

func handleMessage() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte("Broadcast :"+msg))
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
func main() {
	http.HandleFunc("/", handleMultiConnections)
	go handleMessage()
	fmt.Println(" server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
