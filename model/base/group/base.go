package group

import (
	"ApscIM/model/common"
	"context"
)

type Group struct {
	GroupID     int16  `json:"group_id"`
	GroupNo     int16  `json:"group_no"`
	GroupName   string `json:"group_name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	GroupType   int16  `json:"group_type"`
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

type ActionGroupInterface interface {

	/* Base */

	CreateGroup(ctx context.Context, group Group) (err error)
	UpdateGroupInfo(ctx context.Context, groupID int16) (err error)
	GetGroupInfo(ctx context.Context, groupID int16) (group Group, err error)

	/* Member */

	GetGroupMembers(ctx context.Context, groupID int16) (members []MemberGroup, err error)
	AddGroupMembers(ctx context.Context, userIDs []int16) (err error)

	/* Setting */

	GetGroupSetting(ctx context.Context, groupID int16) (setting SettingGroup, err error)
	UpdateGroupSetting(ctx context.Context, setting SettingGroup) (err error)
}
