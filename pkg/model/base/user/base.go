package user

import (
	"ApscIM/pkg/model/common"
	"context"
)

type User struct {
	ID          int32  `json:"id"`
	UserID      int32  `json:"user_id,omitempty"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	Type        int    `json:"type"`
	Description string `json:"description"`
	Agent       bool   `json:"agent"`
	common.BaseTime
}

func (u *User) TableName() string {
	return "apsc_im_users"
}

type SettingUser struct {
	UserID                  int32 `json:"user_id"`
	AllowSameChannelAdd     bool  `json:"allow_same_channel_add"`
	AllowNewMessageNotify   bool  `json:"allow_new_message_notify"`
	QuitChannelClearHistory bool  `json:"quit_channel_clear_history"`
}

func (s *SettingUser) TableName() string {
	return "apsc_im_user_settings"
}

type ActionUserInterface interface {

	/* Base */

	CreateUser(ctx context.Context, user User) (err error)
	UpdateUser(ctx context.Context, user User) (err error)
	GetUsers(ctx context.Context, user User) (users *[]User, err error)

	/* Setting */

	CreateUserSetting(ctx context.Context, setting SettingUser) (err error)
	UpdateUserSetting(ctx context.Context, setting SettingUser) (err error)
	GetUserSetting(ctx context.Context, userID int32) (setting *SettingUser, err error)
}
