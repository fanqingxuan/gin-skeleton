package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func generateRequestId(ctx *gin.Context) {

	requestId := ksuid.New().String()
	ctx.Set("requestId", requestId)
}
