package req

import (
	"ApscIM/pkg/middleware/http"
	"ApscIM/pkg/model/common"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func CheckBindReq[C, R any](ctx *gin.Context, req *R) (*R, error) {
	if err := ctx.BindJSON(&req); err != nil {
		return nil, err
	}
	//if errCheck != nil {
	//	return nil, errCheck
	//}
	return req, nil
}

func WrapResp(ctx *gin.Context, status int, code int, message string, messageID string, success bool, data interface{}) {
	value, ok := ctx.Get("accept-lang")
	var resp interface{}
	var respMessage string
	if ok {
		str, strOk := value.(string)
		if !strOk {
			respMessage = message
		} else {
			fmt.Println(str)
			translateMsg := http.LocalTranslate(str, messageID)
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

func RegisterValidator(tagArr []string, funcArr []validator.Func) error {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("create validator is failed")
	}
	for index, tag := range tagArr {
		fmt.Println(index, tag)
		validate.RegisterValidation(tag, funcArr[index])
	}
	return nil
}
