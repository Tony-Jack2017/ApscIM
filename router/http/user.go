package http

import (
	user2 "ApscIM/api"
	"github.com/gin-gonic/gin"
)

func UserRoutes(userChannel *gin.RouterGroup) {
	userChannel.POST("/login", user2.UserLogin)
}
