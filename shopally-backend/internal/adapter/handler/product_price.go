package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/pkg/domain"
)

func ProductPrice(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, domain.Response{
			Data:  nil,
			Error: "Product ID is required",
		})
		return
	}

	// In a real implementation, fetch the price from a reddis cache or database

	price, exists := productPrices[id]
	if !exists {
		c.JSON(http.StatusNotFound, domain.Response{
			Data:  nil,
			Error: "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data: map[string]interface{}{
			"product_id": id,
			"price":      price,
			"currency":   "USD",
		},
		Error: nil,
	})
}
