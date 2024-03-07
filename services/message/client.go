package message

import "google.golang.org/grpc"

type RpcClientMessage struct {
	conn *grpc.ClientConnInterface
}
