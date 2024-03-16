package api

import (
	"ApscIM/services/user"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	Client *user.RpcClientUser
}

// Base

func (u *UserApi) UserRegister(ctx *gin.Context) {}

func (u *UserApi) UserLogin(ctx *gin.Context) {

}

func (u *UserApi) GetUserInfo(ctx *gin.Context) {}

func (u *UserApi) GetUserList(ctx *gin.Context) {

}

func (u *UserApi) UpdateUser(ctx *gin.Context) {}

// Setting

func (u *UserApi) SetUserSetting(ctx *gin.Context) {}

func (u *UserApi) GetUserSetting(ctx *gin.Context) {}
