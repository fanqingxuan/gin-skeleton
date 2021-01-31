package middleware

import (
	"gin-demo/core/logger"

	"github.com/gin-gonic/gin"
)

func LoadMiddleware(r *gin.Engine) {
	r.Use(generateRequestId)
	r.Use(accessLog(logger.RequestLogger))

	r.Use(recoveryLog)

}
