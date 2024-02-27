package http

import "github.com/gin-gonic/gin"

func MessageRoutes(messageGroup *gin.RouterGroup) {
	messageGroup.GET("/getMessageList")
}
