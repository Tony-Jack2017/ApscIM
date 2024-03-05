package http

import (
	"ApscIM/api/admin/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(adminChannel *gin.RouterChannel) {
	adminChannel.POST("/login", admin.Login)
}
