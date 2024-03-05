package http

import (
	"ApscIM/api/common"
	"github.com/gin-gonic/gin"
)

func CommonRoutes(commonChannel *gin.RouterChannel) {
	commonChannel.POST("/send/codd", common.SendCode)
}
