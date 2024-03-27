package account

import (
	"ApscIM/pkg/tools/req"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type LoginReq struct {
	AccountType string `json:"account_type" binding:"required"`
	IsAdmin     bool   `json:"is_admin" binding:"required"`
	Account     string `json:"account" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"re_password" binding:"eqfield=Password"`
}

type CreateAccountReq struct {
	AccountType string `json:"account_type" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"re_password" binding:"eqfield=Password"`
	Email       string `json:"email"  binding:"email"`
	Country     int    `json:"region"  binding:"act_vil"`
	Mobile      string `json:"mobile"  binding:"act_vil"`
}

type SetAccountSettingReq struct {
	AccountID               int32 `json:"account_id"`
	AllowLoginInOtherDevice bool  `json:"allow_login_in_other_device"`
	AllowLogDeviceOfLogin   bool  `json:"allow_log_device_of_login"`
}

func NewActValidator() error {
	if err := req.RegisterValidator(
		[]string{"act_vil", "act_set_vil"},
		[]validator.Func{ValidatorAccount, ValidatorAccountSetting},
	); err != nil {
		return err
	} else {
		return nil
	}
}

func ValidatorAccount(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
	return matched
}

func ValidatorAccountSetting(fl validator.FieldLevel) bool {
	matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
	return matched
}
