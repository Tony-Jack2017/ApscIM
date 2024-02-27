package websocket

import (
	"encoding/json"
	"fmt"
)

const (
	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

type WsSendMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WsReadMessage struct {
	MessageID int32       `json:"message_id"`
	CreatedAt int32       `json:"created_at"`
	Content   interface{} `json:"content"`
}

func (resp *WsSendMessage) MarshalToJson() (string, error) {
	v, err := json.Marshal(*resp)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func (resp *WsReadMessage) Unmarshal(str string) error {
	err := json.Unmarshal([]byte(str), resp)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WsConnection) ReadMessage() {
	for {
		msgType, message, err := ws.conn.ReadMessage()
		if err != nil {
			fmt.Println("read the message occur the error: ", err)
			return
		}
		switch msgType {
		case TextMessage:
			ws.HandleMessage(msgType, message)
		case BinaryMessage:

		case PongMessage:

		case PingMessage:

		}
	}
}

func (ws *WsConnection) WriteMessage(msgType int, code int, message string, data interface{}) {
	resp := &WsSendMessage{
		Code:    code,
		Message: message,
		Data:    data,
	}
	switch msgType {
	case TextMessage:
		str, err := resp.MarshalToJson()
		if err != nil {
			ws.conn.WriteMessage(msgType, []byte("Something bad happen in server, please contact to CSP"))
		} else {
			ws.conn.WriteMessage(msgType, []byte(str))
		}
		break
	case BinaryMessage:
		str, err := json.Marshal(data)
		if err != nil {
			ws.conn.WriteMessage(msgType, []byte("Something bad happen in server, please contact to CSP"))
		} else {
			ws.conn.WriteMessage(msgType, str)
		}
		break
	case CloseMessage:
		ws.conn.WriteMessage(msgType, []byte("close control message"))
		break
	case PingMessage:
		ws.conn.WriteMessage(msgType, []byte("ping control message"))
		break
	case PongMessage:
		ws.conn.WriteMessage(msgType, []byte("pong control message"))
		break
	}
}

func (ws *WsConnection) HandleMessage(msgType int, message []byte) error {
	req := &WsReadMessage{}
	switch msgType {
	case TextMessage:
		if err := req.Unmarshal(string(message)); err != nil {
			return err
		}
	case BinaryMessage:
	}
	fmt.Println(req)
	return nil
}
