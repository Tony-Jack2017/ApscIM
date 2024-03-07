package user

import (
	"ApscIM/pkg/model/base/account"
	"ApscIM/pkg/model/base/user"
	"ApscIM/services/channel"
	"ApscIM/services/friend"
	"context"
)

type userServer struct {
	accountDatabase account.DatabaseAccount
	userDatabase    user.DatabaseUser
	friendClient    *friend.RpcClientFriend
	channelClient   *channel.RpcClientChannel
}

func Start() {
	user.NewUserDatabase()
}

/* Base */

func (s *userServer) UserRegister(ctx context.Context) {
}

func (s *userServer) SearchDesignateUsers(ctx context.Context) {
}

func (s *userServer) UpdateUserInfo(ctx context.Context) {
}

func (s *userServer) GetUsersByIDs(ctx context.Context) {
}

func (s *userServer) CheckUserAccount(ctx context.Context) {
}

/* Setting */

func (s *userServer) SetUserSetting(ctx context.Context) {
}
