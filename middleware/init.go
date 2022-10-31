package middleware

import (
	"gin-skeleton/svc"

	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddlewares(r *gin.Engine) {

	r.Use(traceId())
	r.Use(requestLog(svc.NewLog("request", "info", svc.RequestLogType)))
	r.Use(recoverLog(svc.NewLog("panic", "info", svc.PanicLogType), svc.NewResponse()))
}
