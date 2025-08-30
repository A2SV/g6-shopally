package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(l *slog.Logger) gin.HandlerFunc {
	if l == nil {
		l = slog.Default()
	}
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start)
		l.Info(
			"http_request",
			"method", method,
			"path", path,
			"status", status,
			"duration", duration.String(),
			"ip", ip,
		)
	}
}
