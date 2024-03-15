package req

import (
	"ApscIM/pkg/middleware/translate"
	"fmt"
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

func WrapResp(ctx *gin.Context, code int, message string, data interface{}) {
	value, ok := ctx.Get("accept-lang")
	if ok {
		str, strOk := value.(string)
		if !strOk {
			ctx.JSON(code, message)
			return
		}
		translateMsg := translate.LocalTranslate(str, message)
		fmt.Println(translateMsg)
		ctx.JSON(code, translateMsg)
		return
	}
	ctx.JSON(code, message)
}
