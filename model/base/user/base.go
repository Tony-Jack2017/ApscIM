package user

type User struct {
	ID          int16  `json:"id"`
	UserID      int16  `json:"user_id,omitempty"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	Type        int    `json:"type"`
	Description string `json:"description"`
	Agent       bool   `json:"agent"`
}

type SettingUser struct {
	AllowSameChannelAdd bool `json:"allow_same_channel_add"`
}

func (u *User) TableName() string {
	return "apsc_im_users"
}
