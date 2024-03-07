package http

import "github.com/gin-gonic/gin"

func MessageRoutes(messageChannel *gin.RouterGroup) {
	messageChannel.GET("/getMessageList")
}
