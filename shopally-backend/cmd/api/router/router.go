package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/cmd/api/middleware"
	"github.com/shopally-ai/internal/adapter/handler"
	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/pkg/domain"
)

func SetupRouter(cfg *config.Config, limiter *middleware.RateLimiter) *gin.Engine {
	router := gin.Default()

	version1 := router.Group("/api/v1")

	// Health checker
	version1.GET("/health", handler.Health)

	//public
	version1.GET("/search", handler.SearchFunc())

	// private
	limitedRouter := version1.Group("")
	limitedRouter.Use(limiter.Middleware())
	{
		limitedRouter.GET("/limited", func(c *gin.Context) {
			c.JSON(http.StatusOK, domain.Response{Data: map[string]interface{}{"message": "limited message"}})
		})
	}
	return router
}
