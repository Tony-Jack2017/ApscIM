package http

import (
	"ApscIM/api/common"
	"github.com/gin-gonic/gin"
)

func CommonRoutes(commonGroup *gin.RouterGroup) {
	commonGroup.POST("/send/codd", common.SendCode)
}
