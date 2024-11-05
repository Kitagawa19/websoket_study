// まず片方のタブでconsoleに入力してからもう片方のタブでconsoleに入力すると、見られる
package main

import ( 
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize : 1024,
	WriteBufferSize : 1024,
	CheckOrigin : func(req  *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade Error:", err)
		return
	}
	defer ws.Close()

	// 接続後にサーバーからメッセージを送信
	err = ws.WriteMessage(websocket.TextMessage, []byte("Welcome to WebSocket server!"))
	if err != nil {
		log.Println("Write Error:", err)
		return
	}

	for {
		// クライアントからのメッセージを読み取る
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		log.Println("Received:", string(msg))
		err = ws.WriteMessage(websocket.TextMessage, []byte("Echo: "+string(msg)))
		if err != nil {
			log.Println("Write Error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/", handleConnections)
	fmt.Println(" server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}