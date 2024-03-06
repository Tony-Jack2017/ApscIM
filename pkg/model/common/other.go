package common

type RespBase struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RespData struct {
	RespBase
	Data interface{} `json:"data"`
}
