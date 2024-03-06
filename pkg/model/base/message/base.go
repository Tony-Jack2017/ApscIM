package message

import (
	"ApscIM/pkg/model/common"
	"context"
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

func (m *Message) TableName() string {
	return "apsc_messages"
}

type SqlMessage struct {
	base Message
}

type SqlMessageInterface interface {

	/* Base */

	CreateMessage(ctx context.Context, message Message) (err error)
	UpdateMessage(ctx context.Context, message Message) (err error)
	GetMessages(ctx context.Context, message Message) (messages *[]Message, err error)
}

type RdsMessageInterface interface {
}

type MgoMessageInterface interface {
}
