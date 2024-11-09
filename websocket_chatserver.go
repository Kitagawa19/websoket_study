package main

import (
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
var broadcast = make(chan Message)

type Message struct {
    Type     string `json:"type"`
    Username string `json:"username"`
    Text     string `json:"text"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade Error:", err)
        return
    }
    defer ws.Close()

    // クライアントをマップに追加
    clients[ws] = true
    log.Println("New client connected. Total clients:", len(clients))

    // クライアントの接続が閉じられた場合、マップから削除
    defer func() {
        delete(clients, ws)
        log.Println("Client disconnected. Remaining clients:", len(clients))
    }()

    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Println("Read Error:", err)
            break
        }

        // メッセージを受け取ったことをログに記録
        log.Printf("Received message from", msg.Username, msg.Text)

        // メッセージをブロードキャストする
        broadcast <- msg
    }
}

func handleMessage() {
    for {
        msg := <-broadcast
        log.Printf("Broadcasting message from", msg.Username, msg.Text)

        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("Write Error: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
    http.HandleFunc("/", handleConnections)
    go handleMessage()
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}