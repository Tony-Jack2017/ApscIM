package http

import (
	"ApscIM/api"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(adminApi *api.AdminApi, adminGroup *gin.RouterGroup) {
	adminGroup.POST("/login", adminApi.AdminLogin)
	adminGroup.POST("/register", adminApi.AdminRegister)
	adminGroup.PUT("/update", adminApi.UpdateAdminInfo)
	adminGroup.GET("/list", adminApi.GetAdminList)

	// setting
	adminGroup.PUT("/setting/set", adminApi.SetAdminSetting)
}
