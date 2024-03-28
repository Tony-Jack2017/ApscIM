package api

import (
	"ApscIM/pkg/middleware/http"
	"ApscIM/pkg/tools/req"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	http.Init()

	r.Use(http.CorsMiddleware()).Use(http.TransMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		req.WrapResp(c, 200, 101, "This is message", "PersonCats", true, nil)
	})

	r.Run("localhost:8080")
}
