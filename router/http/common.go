package http

import (
	"ApscIM/api"
	"github.com/gin-gonic/gin"
)

func CommonRoutes(commonChannel *gin.RouterGroup) {
	commonChannel.POST("/send/codd", api.SendCode)
}
