package admin

import (
	"ApscIM/pkg/model/common"
	"context"
)

type Admin struct {
	AdminID   int32  `json:"admin_id,omitempty"`
	AdminNo   int32  `json:"admin_no"`
	AdminName string `json:"admin_name"`
	Avatar    string `json:"avatar"`
	common.BaseTime
}

type SettingAdmin struct {
	AllowLoginOtherDevice bool `json:"allow_login_other_device"`
}

func (a *Admin) TableName() string {
	return "apsc_im_admins"
}

type SqlAdmin struct {
	base    Admin
	setting SettingAdmin
}

type ActionAdminInterface interface {

	/* Base */

	CreateAdmin(ctx context.Context, admin Admin) (err error)
	GetAdmins(ctx context.Context, condition Admin) (admins *[]Admin, err error)
	UpdateAdmin(ctx context.Context, new Admin) (err error)

	/* Setting */

	CreateAdminSetting(ctx context.Context, admin SettingAdmin) (err error)
	UpdateAdminSetting(ctx context.Context, setting SettingAdmin) (err error)
	GetAdminSetting(ctx context.Context, adminID int32) (setting *SettingAdmin, err error)
}

func CreateAdmin(ctx context.Context, admin Admin) (err error) {
	return nil
}

func GetAdmins(ctx context.Context, condition Admin) (admins *[]Admin, err error) {
	return nil, nil
}

func UpdateAdmin(ctx context.Context, new Admin) (err error) {
	return nil
}

func CreateAdminSetting(ctx context.Context, admin SettingAdmin) (err error) {
	return nil
}

func UpdateAdminSetting(ctx context.Context, setting SettingAdmin) (err error) {
	return nil
}

func GetAdminSetting(ctx context.Context, adminID int32) (setting *SettingAdmin, err error) {
	return nil, nil
}
