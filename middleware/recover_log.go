package middleware

import (
	"errors"
	"fmt"
	"gin-skeleton/common/response"
	"gin-skeleton/svc/logx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				logx.WithContext(ctx).Error(fmt.Sprintf("%+v", err))
				ctx.JSON(http.StatusInternalServerError,
					response.NewCodeError(ctx, http.StatusInternalServerError, errors.New("服务器内部错误")))
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
