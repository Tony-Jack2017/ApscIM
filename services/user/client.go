package user

import "google.golang.org/grpc"

type RpcUserClient struct {
	conn grpc.ClientConnInterface
}
