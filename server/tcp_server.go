package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)
// エコーサーバーを作成する
func tcp_server() {
	listner,err := net.Listen("tcp",":8080") // tcp通信のリスナーを作成
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listner.Close()
	fmt.Println(" Server is listening on port 8080")
	for {
		conn,err := listner.Accept() // クライアントからの接続を待機
		if err != nil {
			fmt.Println(err);
			continue
		}
		fmt.Println("client connected")
		go handleConnection(conn) // クライアントごとに処理を並行して（goroutineを作成して）接続を処理
	}
}

func handleConnection(conn net.Conn) {
	defer  conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message,err :=  reader.ReadString('\n') // クライアントからのメッセージを読み込む
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(message)
		conn.Write([]byte("Echo :" + strings.TrimSpace(message) +"\n"))
	}
}
