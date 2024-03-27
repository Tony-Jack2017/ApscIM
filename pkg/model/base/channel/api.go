package channel

import (
	"ApscIM/pkg/model/common"
	"ApscIM/pkg/tools/req"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type GetChannelListReq struct {
	MemberID      int32             `json:"member_id"`
	ChannelStatus int               `json:"channel_status"`
	ChannelType   int               `json:"channel_type"`
	PageInfo      common.Pagination `json:"page_info"`
}

type CreateChannelReq struct {
	ChannelName  string  `json:"channel_name" binding:"required"`
	ChannelCover string  `json:"channel_cover"`
	ChannelType  int     `json:"channel_type" binding:"required"`
	Description  string  `json:"description"`
	MemberIDS    []int32 `json:"member_ids"`
}

type SetChannelSettingReq struct {
	ChannelID              int32 `json:"channel_id" binding:"required"`
	JoinNeedVerify         bool  `json:"join_need_verify"`
	AllowInvite            bool  `json:"allow_invite"`
	AllowMemberApplyFriend bool  `json:"allow_member_apply_friend"`
	MaxMemberNumber        int   `json:"max_member_number"`
}

func NewChanValidator() error {
	if err := req.RegisterValidator(
		[]string{"chan_vil", "chan_mem_vil", "chan_set_vil"},
		[]validator.Func{ValidatorChannel, ValidatorChannelMember, ValidatorChannelSetting},
	); err != nil {
		return err
	} else {
		return nil
	}
}

func ValidatorChannel(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
	return matched
}

func ValidatorChannelMember(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
	return matched
}

func ValidatorChannelSetting(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
	return matched
}
