package redis

import (
	"ApscIM/pkg/model/base/channel"
	"github.com/redis/go-redis/v9"
)

type ChannelRds struct {
	rdb redis.UniversalClient
}

func NewChannelRds() (channel.RdsChannelInterface, error) {
	return &ChannelRds{}, nil
}

func (c *ChannelRds) AddChannel(channelID int32, userIds []int32) (err error) {
	return nil
}
func (c *ChannelRds) SetChannel(channelID int32) (err error) {
	return nil
}
