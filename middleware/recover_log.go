package middleware

import (
	"errors"
	"fmt"
	"gin-skeleton/common/response"
	"gin-skeleton/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func recoverLog(logger *svc.Log) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				tx_logger := logger.WithContext(ctx)
				tx_logger.Error("", fmt.Sprintf("%+v", err))
				ctx.JSON(http.StatusInternalServerError,
					response.NewCodeError(ctx, http.StatusInternalServerError, errors.New("服务器内部错误")))
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
