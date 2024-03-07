package message

import (
	"ApscIM/pkg/model/base/message"
	"ApscIM/services/channel"
	"ApscIM/services/friend"
	"ApscIM/services/user"
	"context"
)

type messageServer struct {
	messageDatabase message.DatabaseMessage
	userClient      *user.RpcUserClient
	friendClient    *friend.RpcClientFriend
	channelClient   *channel.RpcClientChannel
}

func Start() {
	message.NewMessageDatabase()
}

/* Base */

func (s *messageServer) SendMessage(ctx context.Context) {
}

func (s *messageServer) GetDesignateMessages(ctx context.Context) {
}

func (s *messageServer) InvokeMessage(ctx context.Context) {
}
