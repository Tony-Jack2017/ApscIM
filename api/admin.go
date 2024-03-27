package api

import (
	"ApscIM/services/user"
	"github.com/gin-gonic/gin"
)

type AdminApi struct {
	client *user.RpcClientUser
}

func NewAdminApi() *AdminApi {
	return &AdminApi{
		client: user.NewRpcClientUser(),
	}
}

func AdminRegister(ctx *gin.Context) {
}

func AdminLogin(context *gin.Context) {
}
