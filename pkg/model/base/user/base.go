package user

import (
	"ApscIM/pkg/implementation/mysql"
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
	AccountID   int32  `json:"account_id"`
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

type SqlUserInterface interface {

	/* Base */

	CreateUserSql(ctx context.Context, user *User) (err error)
	UpdateUserSql(ctx context.Context, user *User) (err error)
	GetUsersSql(ctx context.Context, user *User) (users *[]User, err error)

	/* Setting */

	CreateUserSettingSql(ctx context.Context, setting SettingUser) (err error)
	UpdateUserSettingSql(ctx context.Context, setting SettingUser) (err error)
	GetUserSettingSql(ctx context.Context, userID int32) (setting *SettingUser, err error)
}

type RdsUserInterface interface {
}

type MgoUserInterface interface {
	/* Base */

	CreateUserMgo(ctx context.Context, user *User) (err error)
	UpdateUserMgo(ctx context.Context, user *User) (err error)
	GetUsersMgo(ctx context.Context, user *User) (users *[]User, err error)

	/* Setting */

	CreateUserSettingMgo(ctx context.Context, setting SettingUser) (err error)
	UpdateUserSettingMgo(ctx context.Context, setting SettingUser) (err error)
	GetUserSettingMgo(ctx context.Context, userID int32) (setting *SettingUser, err error)
}

type DatabaseUser struct {
	mysql SqlUserInterface
	redis RdsUserInterface
	mongo MgoUserInterface
}

func NewUserDatabase() *DatabaseUser {

	sql, err := mysql.NewUserSql()
	if err != nil {
		panic("connect mysql failed")
	}

	return &DatabaseUser{
		mysql: sql,
		redis: nil,
		mongo: nil,
	}
}
