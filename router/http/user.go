package http

import (
	user2 "ApscIM/api/app/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(userGroup *gin.RouterGroup) {
	userGroup.POST("/login", user2.Login)
}
