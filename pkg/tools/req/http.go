package resp

import (
	"github.com/gin-gonic/gin"
)

func checkBindReq[C, R any](ctx *gin.Context, req *R) (*R, error) {
	if err := ctx.BindJSON(&req); err != nil {
		return nil, err
	}

	errCheck :=

	if errCheck != nil {
		return nil, errCheck
	}
	return req, nil
}
