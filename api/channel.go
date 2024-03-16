package api

import (
	"ApscIM/services/channel"
	"github.com/gin-gonic/gin"
)

type ChannelApi struct {
	Client *channel.RpcClientChannel
}

func CreateChannel(context *gin.Context) {

}

func GetChannelByUser(context *gin.Context) {
}

func UpdateChannelInfo(context *gin.Context) {
}

func SetChannelSetting(context *gin.Context) {
}
