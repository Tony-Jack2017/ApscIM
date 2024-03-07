package channel

import "google.golang.org/grpc"

type RpcClientChannel struct {
	conn grpc.ClientConnInterface
}

func NewRpcClientChannel() RpcClientChannel {
	return RpcClientChannel{}
}
