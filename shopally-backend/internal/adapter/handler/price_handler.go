package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/pkg/util"
)

// PriceHandler exposes endpoints around product pricing.
type PriceHandler struct {
	svc *util.PriceService
}

// NewPriceHandler creates a new PriceHandler.
func NewPriceHandler(svc *util.PriceService) *PriceHandler {
	return &PriceHandler{svc: svc}
}

// GetPrice checks AliExpress for an updated price for a product.
// GET /product/:id/price?current=10.5
func (h *PriceHandler) GetPrice(c *gin.Context) {
	id := c.Param("id")
	currentStr := c.DefaultQuery("current", "0")
	current, _ := strconv.ParseFloat(currentStr, 64)

	updated, changed, err := h.svc.UpdatePriceIfChanged(c.Request.Context(), id, current)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"updated_price": updated, "changed": changed}})
}
