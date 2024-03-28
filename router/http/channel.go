package http

import (
	"ApscIM/api"
	"github.com/gin-gonic/gin"
)

func ChannelRoutes(chanApi *api.ChannelApi, channelGroup *gin.RouterGroup) {
	channelGroup.POST("/create", chanApi.CreateChannel)
	channelGroup.PUT("/update", chanApi.UpdateChannelInfo)
	channelGroup.GET("/list", chanApi.GetChannelByUser)

	// setting
	channelGroup.PUT("/setting/set", chanApi.SetChannelSetting)
}
