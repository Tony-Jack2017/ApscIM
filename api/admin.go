package api

import (
	"ApscIM/pkg/model/base/admin"
	"ApscIM/services/user"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type AdminApi struct {
	client *user.RpcClientUser
	val    *validator.Validate
}

func NewAdminApi(client *user.RpcClientUser) AdminApi {
	var api AdminApi
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("admin_validator", admin.ValidatorAdmin)
		val.RegisterValidation("admin_setting_validator", admin.ValidatorAdminSetting)
	}

	api.client = client
	api.val = val
	return api
}

func AdminRegister(ctx context.Context) {}

func AdminLogin(context *gin.Context) {
}
