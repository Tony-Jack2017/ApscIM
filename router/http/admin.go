package http

import (
	"ApscIM/api/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(adminChannel *gin.RouterGroup) {
	adminChannel.POST("/login", admin.Login)
}
