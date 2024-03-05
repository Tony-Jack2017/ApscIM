package http

import "github.com/gin-gonic/gin"

func Init() {

}

func RegisterRoutesCommon(router *gin.RouterChannel) {
	CommonRoutes(router)
}

func RegisterRoutesFront(router *gin.RouterChannel) {
	message := router.Channel("/message")
	user := router.Channel("/user")
	MessageRoutes(message)
	UserRoutes(user)
}

func RegisterRoutesBack(router *gin.RouterChannel) {
	message := router.Channel("/message")
	user := router.Channel("/user")
	MessageRoutes(message)
	UserRoutes(user)
}

func Run() {
	r := gin.Default()
	v1 := r.Channel("/api/v1")
	back := v1.Channel("/back")
	front := v1.Channel("/front")
	common := v1.Channel("/common")
	RegisterRoutesFront(front)
	RegisterRoutesBack(back)
	RegisterRoutesCommon(common)
}
