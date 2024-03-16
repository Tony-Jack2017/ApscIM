package api

import (
	"ApscIM/services/message"
	"github.com/gin-gonic/gin"
)

type MessageApi struct {
	Client *message.RpcClientMessage
}

func (m *MessageApi) GetMessageList(ctx *gin.Context) {}
