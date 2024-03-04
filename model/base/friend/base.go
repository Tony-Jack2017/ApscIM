package friend

import "ApscIM/model/common"

type Friend struct {
	FriendID int16 `json:"friend_id"`
	UID      int16 `json:"uid"`
	Remark   int16 `json:"remark"`
}

type AddFriendApply struct {
	PromoterID  int16  `json:"promoter_id"`
	RecipientID int16  `json:"recipient_id"`
	ApplyStatus int    `json:"apply_status"`
	ExtraInfo   string `json:"extra_info"`
	Description string `json:"description"`
	common.BaseTime
}

type ActionFriendInterface interface {
	AddFriendRequest(promoterId int16, recipientId int16, remark string, extra string) (err error)
	UpdateFriendInfo(friendId int16, uid int16) (err error)
}
