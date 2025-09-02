package handler

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/shopally-ai/pkg/domain"
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

    // Quick local testing: if ?mock=1 is provided, return an inline mock product
    // (no need to change main.go or gateway wiring). This helps when the real
    // gateway is active but you want deterministic mock responses.
    if c.Query("mock") == "1" {
        fxTs := time.Now().UTC()
        p := &domain.Product{
            ID:                "33006951782",
            Title:             "Mock Sample Phone",
            ImageURL:          "https://via.placeholder.com/300",
            AIMatchPercentage: 90,
            Price:             domain.Price{ETB: 0, USD: 15.90, FXTimestamp: fxTs},
            ProductRating:     4.5,
            SellerScore:       90,
            DeliveryEstimate:  "7-15 days",
            Description:       "Inline mock sample product used for price tests.",
            NumberSold:        1234,
            DeeplinkURL:       "https://www.aliexpress.com/item/33006951782.html",
        }
        updated := p.Price.USD
        changed := false
        const eps = 1e-6
        if (updated - current) > eps || (current-updated) > eps {
            changed = true
        }
    c.JSON(http.StatusOK, gin.H{"data": gin.H{"updated_price": updated, "changed": changed, "product": p}})
        return
    }

    updated, changed, err := h.svc.UpdatePriceIfChanged(c.Request.Context(), id, current)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": gin.H{"updated_price": updated, "changed": changed}})
}
