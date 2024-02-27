package websocket

import (
	"ApscIM/pkg/common/constants"
	"net/http"
)

type WsServer struct {
	Port          int
	MaxConnNumber int
}

func (s *WsServer) Run() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/apsc_im", s.ConnHandle)
	server.ListenAndServe()
}

func (s *WsServer) ConnHandle(writer http.ResponseWriter, req *http.Request) {
	conn := &WsConnection{
		readerBufferSize: 1024,
		writeBufferSize:  1024,
	}
	conn.Start(writer, req)
	if conn.closed {
		writer.Write([]byte("connection build failed"))
	} else {
		conn.WriteMessage(TextMessage, constants.ConnOpenedSuc, "connection build success", "")
	}
}
