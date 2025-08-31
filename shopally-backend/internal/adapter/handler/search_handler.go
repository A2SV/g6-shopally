package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/usecase"
)

// SearchHandler handles incoming HTTP requests for the /search endpoint.
type SearchHandler struct {
	uc *usecase.SearchProductsUseCase
}

// NewSearchHandler creates a new SearchHandler with its dependencies.
func NewSearchHandler(uc *usecase.SearchProductsUseCase) *SearchHandler {
	return &SearchHandler{uc: uc}
}

type envelope struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

// Search handles GET /search and returns the envelope with mocked products.
func (h *SearchHandler) Search(c *gin.Context) {
	// Basic required param validation per contract
	q := strings.TrimSpace(c.Query("q"))
	if q == "" {
		c.JSON(http.StatusBadRequest, envelope{Data: nil, Error: map[string]interface{}{
			"code":    "INVALID_INPUT",
			"message": "missing required query parameter: q",
		}})
		return
	}

	// Optional language preference: am | en (default en) change
	lang := strings.ToLower(strings.TrimSpace(c.DefaultQuery("lang", "en")))
	ctx := c.Request.Context()
	if lang == "am" {
		ctx = context.WithValue(ctx, contextkeys.RespLang, "am")
		ctx = context.WithValue(ctx, contextkeys.RespCurrency, "ETB")
	} else {
		ctx = context.WithValue(ctx, contextkeys.RespLang, "en")
		ctx = context.WithValue(ctx, contextkeys.RespCurrency, "USD")
	}

	data, err := h.uc.Search(ctx, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, envelope{Data: nil, Error: map[string]interface{}{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		}})
		return
	}

	c.JSON(http.StatusOK, envelope{Data: data, Error: nil})
}

// RegisterRoutes sets up the routing for the search handler using Gin.
func (h *SearchHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/search", h.Search)
}
