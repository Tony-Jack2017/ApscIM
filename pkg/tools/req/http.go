package req

import (
	"ApscIM/pkg/middleware/translate"
	"ApscIM/pkg/model/common"
	"github.com/gin-gonic/gin"
)

func checkBindReq[C, R any](ctx *gin.Context, req *R) (*R, error) {
	if err := ctx.BindJSON(&req); err != nil {
		return nil, err
	}
	//if errCheck != nil {
	//	return nil, errCheck
	//}
	return req, nil
}

func WrapResp(ctx *gin.Context, status int, code int, message string, success bool, data interface{}) {
	value, ok := ctx.Get("accept-lang")
	var resp interface{}
	var respMessage string
	if ok {
		str, strOk := value.(string)
		if !strOk {
			respMessage = message
		} else {
			translateMsg := translate.LocalTranslate(str, message)
			respMessage = translateMsg
		}
	} else {
		respMessage = message
	}

	if data != nil {
		resp = &common.RespWithData{
			Code:    code,
			Success: success,
			Message: respMessage,
			Data:    data,
		}
	} else {
		resp = &common.RespBase{
			Code:    code,
			Success: success,
			Message: respMessage,
		}
	}
	ctx.JSON(status, resp)
}
