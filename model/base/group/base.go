package group

import "ApscIM/model/common"

type Group struct {
	GroupID     int16  `json:"group_id"`
	GroupNo     int16  `json:"group_no"`
	GroupName   string `json:"group_name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	common.BaseTime
}

type SettingGroup struct {
	MuteStatus int  `json:"mute_status"`
	Invite     bool `json:"invite"`
}

type MemberGroup struct {
	GroupID       int16  `json:"group_id"`
	UserID        int16  `json:"user_id"`
	GroupRole     string `json:"group_role"`
	Remark        string `json:"remark"`
	SetTop        bool   `json:"set_top"`
	HiddenMessage bool   `json:"hidden_message"`
	//todo
	OtherSetting bool `json:"other_setting"`
}

type NotifyGroup struct {
	NotifyID int16  `json:"notify_id"`
	SenderID int16  `json:"sender_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	common.BaseTime
}
