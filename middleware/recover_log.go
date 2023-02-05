package middleware

import (
	"fmt"
	"gin-skeleton/common/responsex"
	"gin-skeleton/svc/logx"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logx.WithContext(ctx).Error(fmt.Sprintf("%+v", err))
				responsex.ServiceUnavailable(ctx)
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
