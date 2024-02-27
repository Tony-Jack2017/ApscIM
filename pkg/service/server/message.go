package server

import (
	"ApscIM/pkg/service"
	"context"
	"google.golang.org/grpc/profiling/proto"
)

type MessageServer struct {
	service.Server
}

func (m *MessageServer) SendMessage(ctx context.Context, req *proto.EnableRequest) (*proto.EnableRequest, error) {
	// todo
	return req, nil
}

func (m *MessageServer) PushMessage(ctx context.Context, req *proto.ProfilingClient) (*proto.ProfilingClient, error) {
	// todo
	return req, nil
}
