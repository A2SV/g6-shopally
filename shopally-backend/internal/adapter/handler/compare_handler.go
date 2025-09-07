package handler

import (
	"context"
	"net/http"

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

// CompareProducts is the Gin handler for POST /compare.
func (h *CompareHandler) CompareProducts(c *gin.Context) {
	// Require Accept-Language
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Missing required header: Accept-Language",
			},
		})
		return
	}

	var requestBody struct {
		Products []*domain.Product `json:"products"`
	}

	// Parse JSON body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Invalid request body. Ensure it is valid JSON.",
			},
		})
		return
	}

	// Validate number of products
	if len(requestBody.Products) < 2 || len(requestBody.Products) > 4 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INVALID_INPUT",
				"message": "Request body must contain a 'products' array with 2 to 4 product objects.",
			},
		})
		return
	}

	// Attach language to context (support 'am' for Amharic, else default to 'en')
	langCode := lang
	if len(langCode) > 2 {
		langCode = langCode[:2]
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, contextkeys.RespLang, langCode)

	// Execute use case
	comparisonResult, err := h.compareUseCase.Execute(ctx, requestBody.Products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"error": gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "An error occurred while comparing products.",
			},
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{
		"data":  comparisonResult,
		"error": nil,
	})
}
