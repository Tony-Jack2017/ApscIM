package message

type Message struct {
	MessageID   int16  `json:"message_id"`
	MessageType int    `json:"message_type"`
	Content     string `json:"content"`
	ContentType int16  `json:"content_type"`
	ChannelID   int16  `json:"channel_id"`
	SenderID    int16  `json:"sender_id"`
}

func TableName() string {
	return "apsc_messages"
}
