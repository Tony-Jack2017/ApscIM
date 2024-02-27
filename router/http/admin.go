package http

import (
	"ApscIM/api/admin/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(adminGroup *gin.RouterGroup) {
	adminGroup.POST("/login", admin.Login)
}
