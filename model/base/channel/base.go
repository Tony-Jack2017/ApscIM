package channel

import (
	"ApscIM/model/common"
	"context"
)

type Channel struct {
	ChannelID   int16  `json:"channel_id,omitempty"`
	ChannelName string `json:"channel_name"`
	Remark      string `json:"remark"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	common.BaseTime
}

type SettingChannel struct {
}

type ActionChannelInterface interface {

	/* Base */

	CreateChannel(ctx context.Context, channel Channel)
}
