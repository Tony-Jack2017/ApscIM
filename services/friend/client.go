package friend

import "google.golang.org/grpc"

type RpcClientFriend struct {
	conn *grpc.ClientConnInterface
}
