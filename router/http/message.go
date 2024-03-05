package http

import "github.com/gin-gonic/gin"

func MessageRoutes(messageChannel *gin.RouterChannel) {
	messageChannel.GET("/getMessageList")
}
