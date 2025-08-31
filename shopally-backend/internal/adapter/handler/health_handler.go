package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/pkg/domain"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, domain.Response{
		Data: map[string]interface{}{
			"status": "ok",
			"time":   time.Now(),
		},
		Error: nil,
	})
}
