package websocket

import "context"

type WsClient struct {
	Conn       *WsConnection
	PlatformID int16
	Context    context.Context
}
