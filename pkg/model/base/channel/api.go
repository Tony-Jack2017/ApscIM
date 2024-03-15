package channel

import "github.com/go-playground/validator/v10"

type CreateChannelReq struct {
	ChannelName  string  `json:"channel_name"`
	Description  string  `json:"description"`
	ChannelCover string  `json:"channel_cover"`
	MemberIDS    []int32 `json:"member_ids"`
}

type CreateChannelResp struct {
	Success     bool  `json:"success"`
	ChannelID   int32 `json:"Channel_id"`
	ChannelName int32 `json:"Channel_name"`
}

func StructChannelValidation(sl validator.StructLevel) {
	channel := sl.Current().Interface().(Channel)
	if len(channel.ChannelName) == 0 && len(channel.ChannelName) == 0 {
		sl.ReportError(channel.ChannelName, "FirstName", "channel_name", "channel_name", "")
		sl.ReportError(channel.ChannelName, "LastName", "channel_name", "channel_name", "")
	}
}

func StructMemberChannelValidation(sl validator.StructLevel) {
	channel := sl.Current().Interface().(Channel)
	if len(channel.ChannelName) == 0 && len(channel.ChannelName) == 0 {
		sl.ReportError(channel.ChannelName, "FirstName", "channel_name", "channel_name", "")
		sl.ReportError(channel.ChannelName, "LastName", "channel_name", "channel_name", "")
	}
}

func StructSettingChannelValidation(sl validator.StructLevel) {
	channel := sl.Current().Interface().(Channel)
	if len(channel.ChannelName) == 0 && len(channel.ChannelName) == 0 {
		sl.ReportError(channel.ChannelName, "FirstName", "channel_name", "channel_name", "")
		sl.ReportError(channel.ChannelName, "LastName", "channel_name", "channel_name", "")
	}
}

func ValidatorChannel(v *validator.Validate) {
	v.RegisterStructValidation(StructChannelValidation, Channel{})
	v.RegisterStructValidation(StructChannelValidation, SettingChannel{})
	v.RegisterStructValidation(StructChannelValidation, MemberChannel{})
}
