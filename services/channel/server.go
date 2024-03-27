package channel

import (
	"ApscIM/pkg/model/base/channel"
	"context"
)

type channelServer struct {
	channelDatabase channel.DatabaseChannel
	userClient      *RpcClientChannel
}

func Start() {
	channel.NewChannelDatabase()
}

/* Base */

func (s *channelServer) SearchDesignateChannels(ctx context.Context) {

}

func (s *channelServer) CreateChannel(ctx context.Context) {
}

func (s *channelServer) UpdateChannel(ctx context.Context) {
}

func (s *channelServer) ReleaseNotify(ctx context.Context) {
}

/* Setting */

func (s *channelServer) SetChannelSetting(ctx context.Context) {
}

func (s *channelServer) GetChannelSetting(ctx context.Context) {
}

/* Member */

func (s *channelServer) SearchMember(ctx context.Context) {
}

func (s *channelServer) MemberJoinChannel(ctx context.Context) {
}

func (s *channelServer) MemberExitChannel(ctx context.Context) {
}

func (s *channelServer) InviteMember() {
}

func (s *channelServer) UpdateMemberInfo() {
}
