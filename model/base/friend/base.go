package friend

import "ApscIM/model/common"

type Friend struct {
	FriendID int16 `json:"friend_id"`
}

type AddFriendApply struct {
	PromoterID  int16  `json:"promoter_id"`
	RecipientID int16  `json:"recipient_id"`
	ApplyStatus int    `json:"apply_status"`
	ExtraInfo   string `json:"extra_info"`
	Description string `json:"description"`
	common.BaseTime
}
