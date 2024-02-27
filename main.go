package main

import "ApscIM/pkg/common/websocket"

func main() {
	server := &websocket.WsServer{
		Port:          8080,
		MaxConnNumber: 100,
	}
	server.Run()
}
