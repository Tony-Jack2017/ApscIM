package http

import "github.com/gin-gonic/gin"

func Init() {

}

func RegisterRoutesCommon(router *gin.RouterGroup) {
	CommonRoutes(router)
}

func RegisterRoutesFront(router *gin.RouterGroup) {
	message := router.Group("/message")
	user := router.Group("/user")
	MessageRoutes(message)
	UserRoutes(user)
}

func RegisterRoutesBack(router *gin.RouterGroup) {
	message := router.Group("/message")
	user := router.Group("/user")
	MessageRoutes(message)
	UserRoutes(user)
}

func Run() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	back := v1.Group("/back")
	front := v1.Group("/front")
	common := v1.Group("/common")
	RegisterRoutesFront(front)
	RegisterRoutesBack(back)
	RegisterRoutesCommon(common)
}
