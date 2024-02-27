package message

type Message struct {
	MessageID   int16  `json:"message_id"`
	MessageType int    `json:"message_type"`
	Content     string `json:"content"`
}

func TableName() string {
	return "apsc_messages"
}
