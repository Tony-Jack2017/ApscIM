package common

type RespBase struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RespWithData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	Current int `json:"current" form:"current"`
	Size    int `json:"size" form:"size" `
	Total   int `json:"total"`
}
