package mongo

import (
	"ApscIM/pkg/model/base/channel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChannelMgo struct {
	col *mongo.Collection
}

func NewChannelMgo(db *mongo.Database) (channel.MgoChannelInterface, error) {
	col := db.Collection("channel")
	_, err := col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "user_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, err
	}
	return &ChannelMgo{
		col: col,
	}, nil
}

func (c *ChannelMgo) CreateChannel(ctx context.Context, channel channel.Channel) (err error) {
	return nil
}
func (c *ChannelMgo) UpdateChannel(ctx context.Context, channel channel.Channel) (err error) {
	return nil
}
func (c *ChannelMgo) GetChannelInfo(ctx context.Context, channelID int32) (err error) {
	return nil
}

/* Member */

func (c *ChannelMgo) GetChannelMembers(ctx context.Context, channelID int32) (members *[]channel.MemberChannel, err error) {
	return nil, nil
}
func (c *ChannelMgo) AddChannelMembers(ctx context.Context, userIDs []int32) (err error) {
	return nil
}

/* Setting */

func (c *ChannelMgo) GetChannelSetting(ctx context.Context, channelID int32) (setting *channel.SettingChannel, err error) {
	return nil, nil
}
func (c *ChannelMgo) UpdateChannelSetting(ctx context.Context, setting *channel.SettingChannel) (err error) {
	return nil
}

/* Notify */

func (c *ChannelMgo) CreateNotify(ctx context.Context, notify channel.NotifyChannel) (err error) {
	return nil
}
func (c *ChannelMgo) UpdateNotify(ctx context.Context, notify channel.NotifyChannel) (err error) {
	return nil
}
func (c *ChannelMgo) GetNotifies(ctx context.Context, channelID int32) (notifies *[]channel.NotifyChannel, err error) {
	return nil, nil
}
