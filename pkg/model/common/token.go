package common

type Token struct {
	TokenID   int32 `json:"token_id"`
	Expire    int32 `json:"expire"`
	AccountID int32 `json:"account_id"`
}
