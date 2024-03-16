package user

import "google.golang.org/grpc"

type RpcClientUser struct {
	conn grpc.ClientConnInterface
}
