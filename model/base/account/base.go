package account

import "ApscIM/model/common"

type Account struct {
	AccountID int16  `json:"account_id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Region    int    `json:"region"`
	common.BaseTime
}

func TableName() string {
	return "apsc_im_accounts"
}
