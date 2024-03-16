package user

// Base

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
	AccountID int32 `json:"account_id"`
	UserID    int32 `json:"user_id"`
}

type GetUserInfoResp struct {
	User    `json:"user_info"`
	Setting SettingUser `json:"setting"`
}

type GetUserListReq struct {
	UserIds   []int32 `json:"user_ids"`
	Condition User    `json:"condition"`
}

type GetUserResp struct {
	UserList []User `json:"user_list"`
}

// Setting

type SetSettingReq struct {
}

type SetSettingResp struct {
}

type GetSettingReq struct {
	UserId int32 `json:"user_id"`
}

type GetSettingResp struct {
	Setting SettingUser `json:"setting"`
}
