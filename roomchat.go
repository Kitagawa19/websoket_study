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

    var msg Message
    err = ws.ReadJSON(&msg)
    if err == nil && msg.Type == "connect" {
        // 入室メッセージをブロードキャスト
        log.Printf("%s joined the chat", msg.Username)
        broadcast <- Message{Type: "connect", Username: msg.Username, Text: "has joined the chat"}
    }

    // クライアントの接続が閉じられた場合、マップから削除し退室通知を送信
    defer func() {
        delete(clients, ws)
        log.Printf("%s left the chat", msg.Username)
        broadcast <- Message{Type: "disconnect", Username: msg.Username, Text: "has left the chat"}
    }()

    // メッセージを受信し、ブロードキャスト
    for {
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Println("Read Error:", err)
            break
        }

        // ブロードキャストするのは "message" タイプのみ
        if msg.Type == "message" {
            broadcast <- msg
        }
    }
}

func handleMessage() {
    for {
        msg := <-broadcast
        log.Printf("Broadcasting message from %s: %s", msg.Username, msg.Text)

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