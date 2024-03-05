package user

type UpdateUserInfoReq struct {
	UserID      int32  `json:"user_id"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type UpdateUserInfoResp struct {
	Success bool `json:"success"`
}

type GetUserInfoReq struct {
	UserID int32 `json:"user_id"`
}
