package account

import (
	"ApscIM/pkg/model/common"
	"context"
)

type Account struct {
	ID        int32  `json:"id"`
	AccountID int32  `json:"account_id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Region    int    `json:"region"`
	common.BaseTime
}

func (a *Account) TableName() string {
	return "apsc_im_accounts"
}

type SettingAccount struct {
	ID                      int32 `json:"id"`
	AllowLoginInOtherDevice bool  `json:"allow_login_in_other_device"`
	AllowLogDeviceOfLogin   bool  `json:"allow_log_device_of_login"`
}

func (s *SettingAccount) TableName() string {
	return "apsc_im_account_settings"
}

type SQLAccount struct {
	base    Account
	setting SettingAccount
}

type ActionAccountInterface interface {

	/* Base */

	AddAccount(ctx context.Context, account Account) (err error)
	UpdateAccount(ctx context.Context, account Account) (err error)
	GetAccounts(ctx context.Context, condition map[string]string) (account []Account, err error)

	/* Setting */

	AddAccountSetting(ctx context.Context, setting SettingAccount) (err error)
	GetAccountSetting(ctx context.Context, accountID int32) (setting SettingAccount, err error)
	UpdateAccountSetting(ctx context.Context, setting SettingAccount) (err error)
}

func (sql *SQLAccount) AddAccount() {
}
