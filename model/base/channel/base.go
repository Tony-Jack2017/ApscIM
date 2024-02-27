package channel

import "ApscIM/model/common"

type Channel struct {
	ChannelID int16  `json:"channel_id,omitempty"`
	Remark    string `json:"remark"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	common.BaseTime
}
