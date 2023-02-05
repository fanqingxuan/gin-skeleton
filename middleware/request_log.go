package middleware

import (
	"fmt"
	"gin-skeleton/svc/logx"
	"time"

	"github.com/gin-gonic/gin"
)

func requestLog() gin.HandlerFunc {

	notlogged := []string{}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {

			if raw != "" {
				path = path + "?" + raw
			}
			TimeStamp := time.Now()
			Latency := TimeStamp.Sub(start)
			if Latency > time.Minute {
				Latency = Latency.Truncate(time.Second)
			}
			requestMap := map[string]interface{}{
				"Path":      path,
				"Method":    c.Request.Method,
				"ClientIP":  c.ClientIP(),
				"Latency":   fmt.Sprintf("%s", Latency),
				"Status":    c.Writer.Status(),
				"Proto":     c.Request.Proto,
				"UserAgent": c.Request.UserAgent(),
				"Msg":       c.Errors.ByType(gin.ErrorTypePrivate).String(),
				"Size":      c.Writer.Size(),
			}

			logx.WithContext(c).Info(requestMap)
		}
	}
}
