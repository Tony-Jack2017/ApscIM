package http

import (
	user2 "ApscIM/api/app/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(userChannel *gin.RouterChannel) {
	userChannel.POST("/login", user2.Login)
}
