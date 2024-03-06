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
	AccountID int32  `json:"account_id"`
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

type SqlAdminInterface interface {

	/* Base */

	CreateAdmin(ctx context.Context, admin Admin) (err error)
	GetAdmins(ctx context.Context, condition Admin) (admins *[]Admin, err error)
	UpdateAdmin(ctx context.Context, new Admin) (err error)

	/* Setting */

	CreateAdminSetting(ctx context.Context, admin SettingAdmin) (err error)
	UpdateAdminSetting(ctx context.Context, setting SettingAdmin) (err error)
	GetAdminSetting(ctx context.Context, adminID int32) (setting *SettingAdmin, err error)
}

type RdsAdminInterface interface {
}

type MgoAdminInterface interface {
}
