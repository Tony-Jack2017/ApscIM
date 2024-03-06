package channel

type DatabaseChannel struct {
	mysql SqlChannelInterface
	redis RdsChannelInterface
	mongo MgoChannelInterface
}

func NewChannelDatabase() *DatabaseChannel {
	return &DatabaseChannel{
		mysql: nil,
		redis: nil,
		mongo: nil,
	}
}
