package user

import (
	"ApscIM/pkg/model/common"
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

type SettingUser struct {
	AllowSameChannelAdd     bool `json:"allow_same_channel_add"`
	AllowNewMessageNotify   bool `json:"allow_new_message_notify"`
	QuitChannelClearHistory bool `json:"quit_channel_clear_history"`
}

func (u *User) TableName() string {
	return "apsc_im_users"
}

type ActionUserInterface interface {
}
