package friend

import (
	"ApscIM/pkg/model/common"
	"context"
)

type Friend struct {
	FriendID     int32 `json:"friend_id"`
	UserID       int32 `json:"user_id"`
	Remark       int32 `json:"remark"`
	FriendStatus int   `json:"friend_status"`
	FriendType   int   `json:"friend_type"`
	common.BaseTime
}

func (f *Friend) TableName() string {
	return "apsc_im_friends"
}

type AddFriendApply struct {
	PromoterID  int32  `json:"promoter_id"`
	RecipientID int32  `json:"recipient_id"`
	ApplyStatus int    `json:"apply_status"`
	ExtraInfo   string `json:"extra_info"`
	Description string `json:"description"`
	common.BaseTime
}

func (a *AddFriendApply) TableName() string {
	return "apsc_im_friend_applies"
}

type SettingFriend struct {
}

func (s *SettingFriend) TableName() string {
	return "apsc_im_friend_settings"
}

type SqlFriend struct {
	base    Friend
	apply   AddFriendApply
	setting SettingFriend
}

type SqlFriendInterface interface {

	/* Base */

	CreateFriend(ctx context.Context, friend Friend) (err error)
	UpdateFriend(ctx context.Context, friendId int32, uid int32) (err error)

	/* Apply */

	CreateApply(ctx context.Context, apply AddFriendApply) (err error)
	UpdateApply(ctx context.Context, apply AddFriendApply) (err error)
	GetApplies(ctx context.Context)

	/* Setting */

	AddFriendSetting(ctx context.Context, setting SettingFriend) (err error)
	UpdateFriendSetting(ctx context.Context, setting SettingFriend) (err error)
	GetFriendSetting(ctx context.Context, friendID int32, userID int32) (setting *SettingFriend, err error)
}

type RdsFriendInterface interface {
}

type MgoFriendInterface interface {
}
