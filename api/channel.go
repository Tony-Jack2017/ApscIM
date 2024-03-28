package api

import (
	"ApscIM/services/channel"
	"github.com/gin-gonic/gin"
)

type ChannelApi struct {
	client *channel.RpcClientChannel
}

func NewChannelApi() *ChannelApi {
	client := channel.NewRpcClientChannel()
	return &ChannelApi{
		client: client,
	}
}

func (c *ChannelApi) CreateChannel(ctx *gin.Context) {
}

func (c *ChannelApi) GetChannelByUser(ctx *gin.Context) {
}

func (c *ChannelApi) UpdateChannelInfo(ctx *gin.Context) {
}

// Setting

func (c *ChannelApi) SetChannelSetting(ctx *gin.Context) {
}
