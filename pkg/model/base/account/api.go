package account

type LoginReq struct {
	AccountType string `json:"account_type"`
	Account     string `json:"account"`
	Password    string `json:"password"`
}

type LoginResp struct {
}

type RegisterAccountReq struct {
	AccountType string `json:"account_type"`
	Password    string `json:"password"`
}
