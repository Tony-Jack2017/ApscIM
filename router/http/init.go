package http

import (
	"ApscIM/pkg/middleware/common"
	"ApscIM/pkg/middleware/translate"
	"ApscIM/pkg/model/base/channel"
	"ApscIM/pkg/tools/req"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		channel.ValidatorChannel(v)
	}
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

	translate.Init()

	r.Use(common.CorsMiddleware()).Use(translate.TransMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		req.WrapResp(c, 200, "This is message", "")
	})

	r.Run("localhost:8080")

	// routes
	//v1 := r.Group("/api/v1")
	//back := v1.Group("/back")
	//front := v1.Group("/front")
	//common := v1.Group("/common")
	//RegisterRoutesFront(front)
	//RegisterRoutesBack(back)
	//RegisterRoutesCommon(common)
}
