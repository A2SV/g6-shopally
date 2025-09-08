package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/cmd/api/middleware"
	"github.com/shopally-ai/internal/adapter/handler"
	"github.com/shopally-ai/internal/config"
)

func SetupRouter(cfg *config.Config, limiter *middleware.RateLimiter, searchHandler *handler.SearchHandler, compareHandler *handler.CompareHandler, alertHandler *handler.AlertHandler, priceHandler *handler.PriceHandler) *gin.Engine {
	router := gin.Default()

	version1 := router.Group("/api/v1")
	version1.GET("/health", handler.Health)

	limitedRouter := version1.Group("")
	limitedRouter.Use(limiter.Middleware())
	{
		limitedRouter.POST("/compare", compareHandler.CompareProducts)
		limitedRouter.GET("/search", searchHandler.Search)
	}
	version1.POST("/alerts", alertHandler.CreateAlertHandler)
	version1.GET("/alerts/:id", alertHandler.GetAlertHandler)
	version1.DELETE("/alerts/:id", alertHandler.DeleteAlertHandler)
	version1.GET("/product/:id/price", priceHandler.GetPrice)

	return router
}
