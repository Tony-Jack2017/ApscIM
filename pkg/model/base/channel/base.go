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

func (c *Channel) TableName() string {
	return "apsc_im_channels"
}

type SettingChannel struct {
	ChannelID              int32 `json:"channel_id"`
	JoinNeedVerify         bool  `json:"join_need_verify"`
	AllowInvite            bool  `json:"allow_invite"`
	AllowMemberApplyFriend bool  `json:"allow_member_apply_friend"`
	MaxMemberNumber        int   `json:"max_member_number"`
	common.BaseTime
}

func (s *SettingChannel) TableName() string {
	return "apsc_im_channel_settings"
}

type MemberChannel struct {
	ChannelID     int32  `json:"Channel_id"`
	UserID        int32  `json:"user_id"`
	RoleInChannel int    `json:"role_in_Channel"`
	Nickname      string `json:"nickname"`
	NoticeStatus  int    `json:"notice_status"`
}

func (m *MemberChannel) TableName() string {
	return "apsc_im_channel_members"
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

func (n *NotifyChannel) TableName() string {
	return "apsc_im_channel_notifies"
}

type SqlChannel struct {
	base    Channel
	setting SettingChannel
	member  MemberChannel
	notify  NotifyChannel
}

type SqlChannelInterface interface {

	/* Base */

	CreateChannel(ctx context.Context, channel Channel) (err error)
	UpdateChannelInfo(ctx context.Context, channelID int32) (err error)
	GetChannelInfo(ctx context.Context, channelID int32) (channel *Channel, err error)

	/* Member */

	GetChannelMembers(ctx context.Context, channelID int32) (members *[]MemberChannel, err error)
	AddChannelMembers(ctx context.Context, userIDs []int32) (err error)

	/* Setting */

	GetChannelSetting(ctx context.Context, channelID int32) (setting *SettingChannel, err error)
	UpdateChannelSetting(ctx context.Context, setting *SettingChannel) (err error)

	/* Notify */

	CreateNotify(ctx context.Context, notify NotifyChannel) (err error)
	UpdateNotify(ctx context.Context, notify NotifyChannel) (err error)
	GetNotifies(ctx context.Context, channelID int32) (notifies *[]NotifyChannel, err error)
}

type RdsChannelInterface interface {
}

type MgoChannelInterface interface {

	/* Base */

	CreateChannel(ctx context.Context, channel Channel) (err error)
	UpdateChannel(ctx context.Context, channel Channel) (err error)
	GetChannelInfo(ctx context.Context, channelID int32) (err error)

	/* Member */

	GetChannelMembers(ctx context.Context, channelID int32) (members *[]MemberChannel, err error)
	AddChannelMembers(ctx context.Context, userIDs []int32) (err error)

	/* Setting */

	GetChannelSetting(ctx context.Context, channelID int32) (setting *SettingChannel, err error)
	UpdateChannelSetting(ctx context.Context, setting *SettingChannel) (err error)

	/* Notify */

	CreateNotify(ctx context.Context, notify NotifyChannel) (err error)
	UpdateNotify(ctx context.Context, notify NotifyChannel) (err error)
	GetNotifies(ctx context.Context, channelID int32) (notifies *[]NotifyChannel, err error)
}
