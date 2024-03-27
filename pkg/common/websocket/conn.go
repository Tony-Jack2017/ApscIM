package websocket

import (
	"ApscIM/pkg/common/constants"
	http2 "ApscIM/pkg/middleware/http"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

type WsConnection struct {
	clientID         int32
	userID           int32
	conn             *websocket.Conn
	handshakeTimeout time.Duration
	writeBufferSize  int
	readerBufferSize int
	closed           bool
}

func (ws *WsConnection) ParseArgs(writer http.ResponseWriter, req *http.Request) error {
	query := req.URL.Query()

	clientId := query.Get("client_id")
	tokenStr := query.Get("apsc_token")
	userId := query.Get("user_id")

	if clientId == "" {
		return errors.New("the 'client_id' is null, request params can't be null")
	} else {
		cache, err := strconv.ParseInt(clientId, 10, 16)
		if err != nil {
			return errors.New("the type of 'client_id' must is number")
		} else {
			ws.clientID = int32(cache)
		}
	}

	if userId == "" {
		return errors.New("the 'user_id' is null, request params can't be null")
	} else {
		cache, err := strconv.ParseInt(userId, 10, 16)
		if err != nil {
			writer.WriteHeader(400)
			writer.Write([]byte("the type of 'user_id' must is number !!!"))
			return errors.New("the type of 'user_id' must is number")
		} else {
			ws.clientID = int32(cache)
		}
	}

	if tokenStr == "" {
		return errors.New("the token of connection is null")
	} else {
		_, err := http2.ParseToken(tokenStr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsConnection) Upgrade(writer http.ResponseWriter, req *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:   ws.readerBufferSize,
		WriteBufferSize:  ws.writeBufferSize,
		HandshakeTimeout: ws.handshakeTimeout,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return upgrader.Upgrade(writer, req, nil)
}

func (ws *WsConnection) Start(writer http.ResponseWriter, req *http.Request) {
	//if err := ws.ParseArgs(writer, req); err != nil {
	//	return err
	//}
	conn, err := ws.Upgrade(writer, req)
	if err != nil {
		fmt.Println("upgrade websocket occur error: ", err)
		return
	}
	ws.conn = conn
	go ws.ReadMessage()
}

func (ws *WsConnection) Close(err error) error {
	ws.closed = true
	if err != nil {
		ws.WriteMessage(CloseMessage, constants.ConnClosing,
			"There's a error happening in server, websocket connection is closing...",
			map[string]error{
				"error": err,
			},
		)
	} else {
		ws.WriteMessage(CloseMessage, constants.ConnClosing, "websocket connection is closing...", "")
	}
	return ws.conn.Close()
}
