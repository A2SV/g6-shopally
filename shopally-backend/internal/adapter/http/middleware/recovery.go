package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(l *slog.Logger) gin.HandlerFunc {
	if l == nil {
		l = slog.Default()
	}

	return func (c *gin.Context)  {
		defer func() {
			if r := recover(); r != nil {
				l.Error("panic_recovered", "error", r, "path", c.Request.URL.Path)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
		}()
		c.Next()
	}
}