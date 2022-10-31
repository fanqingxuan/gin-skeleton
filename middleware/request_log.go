package middleware

import (
	"gin-skeleton/svc"

	"github.com/gin-gonic/gin"
)

func requestLog(logger *svc.Log) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		requestInfo := map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"client": c.ClientIP(),
			"query":  c.Request.URL.RawQuery,
			"post":   c.Request.PostForm,
		}
		logger.WithContext(c).Info("request", requestInfo)
		c.Next()
	}
}
