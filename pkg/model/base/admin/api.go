package admin

import (
	"ApscIM/pkg/model/common"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type RegisterAdminReq struct {
	AdminName  string `json:"admin_name" validate:"required,admin_validator"`
	InviteCode int    `json:"invite_code"`
}

type RegisterAdminResp struct {
	common.RespBase
}

func ValidatorAdmin(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
	return matched
}

func ValidatorAdminSetting(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
	return matched
}
