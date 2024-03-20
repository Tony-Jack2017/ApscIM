package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RpcClientUser struct {
	conn grpc.ClientConnInterface
	//Client UserClient
}

func NewRpcClientUser() *RpcClientUser {
	conn, err := grpc.Dial("127.0.0.1:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	return &RpcClientUser{
		conn: conn,
	}
}
