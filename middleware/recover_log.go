package middleware

import (
	"fmt"
	"gin-skeleton/common/errorx"
	"gin-skeleton/common/responsex"
	"gin-skeleton/svc/logx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logx.WithContext(ctx).Error(fmt.Sprintf("%+v", err))
				responsex.New(ctx, errorx.NewCodeError(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable)))
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
