package middleware

import (
	"fmt"
	"gin-skeleton/svc"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

func recoverLog(logger *svc.Log, resp *svc.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				tx_logger := logger.WithContext(ctx)
				tx_logger.Error("response error", err)
				fmt.Println(string(debug.Stack()))
				tx_logger.Error("response error detail", strings.ReplaceAll(string(debug.Stack()), "\\n", "\\r\\n"))
				ctx.JSON(http.StatusInternalServerError, resp.ErrorWithCode(http.StatusInternalServerError, "服务器内部错误"))
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
