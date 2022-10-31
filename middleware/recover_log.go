package middleware

import (
	"fmt"
	"gin-skeleton/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func recoverLog(logger *svc.Log, resp *svc.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				tx_logger := logger.WithContext(ctx)
				tx_logger.Error("", fmt.Sprintf("%+v", err))
				ctx.JSON(http.StatusInternalServerError, resp.ErrorWithCode(http.StatusInternalServerError, "服务器内部错误"))
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
