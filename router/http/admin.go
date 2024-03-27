package http

import (
	"ApscIM/api"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(adminChannel *gin.RouterGroup) {
	adminChannel.POST("/login", api.AdminLogin)
	adminChannel.POST("/register", api.AdminRegister)
}
