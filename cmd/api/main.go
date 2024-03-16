package api

import (
	"ApscIM/pkg/middleware/common"
	"ApscIM/pkg/middleware/translate"
	"ApscIM/pkg/tools/req"
	"github.com/gin-gonic/gin"
)

func Run(svcPort int) {
	r := gin.Default()

	translate.Init()

	r.Use(common.CorsMiddleware()).Use(translate.TransMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		req.WrapResp(c, 200, "This is message", "")
	})

	r.Run("localhost:8080")
}
