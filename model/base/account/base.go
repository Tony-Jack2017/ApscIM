package account

import "ApscIM/model/common"

type Account struct {
	AccountID int16  `json:"account_id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	common.BaseTime
}
