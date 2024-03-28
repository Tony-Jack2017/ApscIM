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

func (a *AdminApi) AdminRegister(ctx *gin.Context) {
}

func (a *AdminApi) AdminLogin(ctx *gin.Context) {
}

func (a *AdminApi) UpdateAdminInfo(ctx *gin.Context) {
}

func (a *AdminApi) GetAdminList(ctx *gin.Context) {
}

// Setting

func (a *AdminApi) SetAdminSetting(ctx *gin.Context) {
}
