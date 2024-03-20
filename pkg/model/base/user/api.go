package user

import "github.com/go-playground/validator/v10"

// Base

type UpdateUserInfoReq struct {
	UserID      int32  `json:"user_id"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}
type GetUserInfoReq struct {
	AccountID int32 `json:"account_id" validate:"required,valid_username"`
	UserID    int32 `json:"user_id"`
}
type GetUserListReq struct {
	UserIds   []int32 `json:"user_ids"`
	Condition User    `json:"condition"`
}

type UpdateUserInfoResp struct {
	Success bool `json:"success"`
}
type GetUserInfoResp struct {
	User    `json:"user_info"`
	Setting SettingUser `json:"setting"`
}
type GetUserResp struct {
	UserList []User `json:"user_list"`
}

// Setting

type SetSettingReq struct {
}
type GetSettingReq struct {
	UserId int32 `json:"user_id"`
}

type SetSettingResp struct {
}
type GetSettingResp struct {
	Setting SettingUser `json:"setting"`
}

func ValidatorUser(fl validator.FieldLevel) {
}

func ValidatorUserSetting() {
}
