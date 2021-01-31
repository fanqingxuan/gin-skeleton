package middleware

import (
	"fmt"
	"gin-demo/controller"
	"gin-demo/core/logger"
	"gin-demo/util"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
)

func recoveryLog(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

			err_message := err
			str := string(debug.Stack())

			switch val := err.(type) {
			case logger.CustomLog:
				err_message = util.TrimmedPath(val.File) + ":" + strconv.Itoa(val.Line) + "\t" + val.Message
				str = ""
			}
			message := fmt.Sprintf("%s\t%s\t%s\t%s\r\n%s",
				ctx.Value("requestId").(string), err_message, ctx.Request.URL.Path, ctx.Request.URL.RawQuery, str)

			logger.PanicLogger.Error(message)
			controller.StatuscoreServerError(ctx)
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}
