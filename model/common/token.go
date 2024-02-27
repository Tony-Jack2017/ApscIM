package common

type Token struct {
	TokenID   int16 `json:"token_id"`
	Expire    int16 `json:"expire"`
	AccountID int16 `json:"account_id"`
}
