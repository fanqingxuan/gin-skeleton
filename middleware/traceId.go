package middleware

import (
	"gin-skeleton/common/utils"

	"github.com/gin-gonic/gin"
)

func traceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := utils.NewUuid()
		ctx.Set("traceId", traceId)
		ctx.Next()
	}
}
