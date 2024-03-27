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

func (c *ChannelApi) CreateChannel(context *gin.Context) {

}

func (c *ChannelApi) GetChannelByUser(context *gin.Context) {
}

func (c *ChannelApi) UpdateChannelInfo(context *gin.Context) {
}

func (c *ChannelApi) SetChannelSetting(context *gin.Context) {
}
