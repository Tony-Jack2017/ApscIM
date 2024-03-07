package http

import (
	"ApscIM/api/common"
	"github.com/gin-gonic/gin"
)

func CommonRoutes(commonChannel *gin.RouterGroup) {
	commonChannel.POST("/send/codd", common.SendCode)
}
