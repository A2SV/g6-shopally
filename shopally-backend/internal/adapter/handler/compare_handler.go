package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"
)

// CompareHandler handles HTTP requests related to product comparison.
type CompareHandler struct {
	compareUseCase usecase.CompareProductsExecutor
}

// NewCompareHandler creates a new instance of CompareHandler.
func NewCompareHandler(uc usecase.CompareProductsExecutor) *CompareHandler {
	return &CompareHandler{
		compareUseCase: uc,
	}
}

func (h *CompareHandler) CompareProducts(c *gin.Context) {
	// Derive a context with timeout from the incoming request
	ctx, cancel := context.WithTimeout(c.Request.Context(), 120*time.Second)
	defer cancel()

	// Extract headers and attach them as values to the context
	deviceID := strings.TrimSpace(c.GetHeader("X-Device-ID"))
	lang := strings.ToLower(strings.TrimSpace(c.GetHeader("Accept-Language")))

	if deviceID == "" || lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Missing required headers: X-Device-ID or Accept-Language",
			},
		})
		return
	}

	ctx = context.WithValue(ctx, contextkeys.DeviceID, deviceID)
	ctx = context.WithValue(ctx, contextkeys.RespLang, lang)

	// Parse request body
	var reqBody struct {
		Products []*domain.Product `json:"products"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Invalid JSON body",
			},
		})
		return
	}

	// Validate product count
	if len(reqBody.Products) < 2 || len(reqBody.Products) > 4 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Products array must contain 2 to 4 items",
			},
		})
		return
	}

	// Pass the timeout-aware, header-enriched context to the use case
	comparisonResult, err := h.compareUseCase.Execute(ctx, reqBody.Products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": err.Error(),
			},
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"data":  comparisonResult,
		"error": nil,
	})
}
