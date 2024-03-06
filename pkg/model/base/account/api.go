package account

import "ApscIM/pkg/model/common"

type LoginReq struct {
	AccountType string `json:"account_type"`
	Account     string `json:"account"`
	Password    string `json:"password"`
}

type LoginResp struct {
	common.RespData
	Token string
}

type RegisterAccountReq struct {
	AccountType string `json:"account_type"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Region      string `json:"region"`
	Mobile      string `json:"mobile"`
}

type RegisterAccountResp struct {
	common.RespBase
}
