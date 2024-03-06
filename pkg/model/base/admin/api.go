package admin

import "ApscIM/pkg/model/common"

type CreateAdminReq struct {
	AdminName string `json:"admin_name"`
}

type CreateAdminResp struct {
	common.RespBase
}
