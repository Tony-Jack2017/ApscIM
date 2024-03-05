package message

import (
	"ApscIM/pkg/model/common"
)

type Message struct {
	MessageID   int32  `json:"message_id"`
	MessageType int    `json:"message_type"`
	Content     string `json:"content"`
	ContentType int32  `json:"content_type"`
	ChannelID   int32  `json:"channel_id"`
	SenderID    int32  `json:"sender_id"`
	common.BaseTime
}

func TableName() string {
	return "apsc_messages"
}

type ActionMessageInterface interface {
}
