package channel

import (
	"ApscIM/pkg/model/common"
	"context"
)

type Channel struct {
	ChannelID   int32  `json:"channel_id"`
	ChannelNo   int32  `json:"channel_no"`
	ChannelName string `json:"channel_name"`
	ChannelType int32  `json:"channel_type"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	InviteLink  string `json:"invite_link"`
	Extra       string `json:"extra"`
	//ChannelPublic int32  `json:"channel_public"`	// waiting
	common.BaseTime
}

type SettingChannel struct {
	ChannelID              int32 `json:"channel_id"`
	MuteStatus             int   `json:"mute_status"`
	JoinNeedVerify         bool  `json:"join_need_verify"`
	AllowInvite            bool  `json:"allow_invite"`
	AllowMemberApplyFriend bool  `json:"allow_member_apply_friend"`
	common.BaseTime
}

type MemberChannel struct {
	ChannelID     int32  `json:"Channel_id"`
	UserID        int32  `json:"user_id"`
	RoleInChannel int    `json:"role_in_Channel"`
	Nickname      string `json:"nickname"`
	NoticeStatus  int    `json:"notice_status"`
}

type NotifyChannel struct {
	ChannelID    int32  `json:"Channel_id"`
	NotifyID     int32  `json:"notify_id"`
	NotifyStatus int32  `json:"notify_status"`
	SenderID     int32  `json:"sender_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	common.BaseTime
}

type ActionChannelInterface interface {

	/* Base */

	CreateChannel(ctx context.Context, channel Channel) (err error)
	UpdateChannelInfo(ctx context.Context, ChannelID int32) (err error)
	GetChannelInfo(ctx context.Context, ChannelID int32) (channel *Channel, err error)

	/* Member */

	GetChannelMembers(ctx context.Context, ChannelID int32) (members *[]MemberChannel, err error)
	AddChannelMembers(ctx context.Context, userIDs []int32) (err error)

	/* Setting */

	GetChannelSetting(ctx context.Context, ChannelID int32) (setting *SettingChannel, err error)
	UpdateChannelSetting(ctx context.Context, setting *SettingChannel) (err error)

	/* Notify */

	CreateNotify(ctx context.Context, notify NotifyChannel) (err error)
	UpdateNotify(ctx context.Context, notify NotifyChannel) (err error)
	GetNotifies(ctx context.Context, ChannelID int32) (notifies *[]NotifyChannel, err error)
}
