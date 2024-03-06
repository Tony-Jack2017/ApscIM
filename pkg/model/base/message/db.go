package message

type DatabaseMessage struct {
	mysql SqlMessageInterface
	redis RdsMessageInterface
	mongo MgoMessageInterface
}

func NewMessageDatabase() {
}
