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

	// Mock product price data
	productPrices := map[string]float64{
		"1": 19.99,
		"2": 29.99,
		"3": 39.99,
	}

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
