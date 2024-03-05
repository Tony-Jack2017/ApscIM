package websocket

import "context"

type WsClient struct {
	Conn       *WsConnection
	PlatformID int32
	Context    context.Context
}
