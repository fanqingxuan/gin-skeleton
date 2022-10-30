package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func traceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := uuid.NewV4()
		ctx.Set("traceId", traceId)
		ctx.Next()
	}
}
