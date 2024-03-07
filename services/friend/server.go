package friend

import (
	"ApscIM/pkg/model/base/channel"
	"ApscIM/pkg/model/base/friend"
	"ApscIM/services/user"
	"context"
)

type friendServer struct {
	friendDatabase friend.DatabaseFriend
	userClient     *user.RpcUserClient
}

func Start() {
	channel.NewChannelDatabase()
}

/* Base */

func (s *friendServer) AddFriendApply(ctx context.Context) {
}

func (s *friendServer) DeleteFriend(ctx context.Context) {
}

/* Setting */

func (s *friendServer) SetFriendSetting() {
}
