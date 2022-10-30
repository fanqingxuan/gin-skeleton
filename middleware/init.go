package middleware

import (
	"gin-skeleton/svc"

	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddlewares(r *gin.Engine) {

	r.Use(traceId())
	r.Use(accessLog(svc.NewLog("access_log", "info")))
	r.Use(recoverLog(svc.NewLog("error_log", "info"), svc.NewResponse()))
}
