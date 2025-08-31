package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

// Package-level injection so the router can register the handler without a
// direct dependency on a handler instance. Main should call InjectSearchUseCase
// during startup with the appropriate usecase.
var searchUseCase *usecase.SearchProductsUseCase

// InjectSearchUseCase provides the usecase to the package-level handler func.
func InjectSearchUseCase(uc *usecase.SearchProductsUseCase) {
	searchUseCase = uc
}

// SearchFunc returns a gin.HandlerFunc that uses the injected usecase to
// perform searches. This allows router.SetupRouter to register the route
// without receiving a handler instance as a parameter.
func SearchFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Basic required param validation per contract
		q := strings.TrimSpace(c.Query("q"))
		if q == "" {
			c.JSON(http.StatusBadRequest, envelope{Data: nil, Error: map[string]interface{}{
				"code":    "INVALID_INPUT",
				"message": "missing required query parameter: q",
			}})
			return
		}

		if searchUseCase == nil {
			c.JSON(http.StatusInternalServerError, envelope{Data: nil, Error: map[string]interface{}{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "search usecase not initialized",
			}})
			return
		}

		data, err := searchUseCase.Search(c.Request.Context(), q)
		if err != nil {
			c.JSON(http.StatusInternalServerError, envelope{Data: nil, Error: map[string]interface{}{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": err.Error(),
			}})
			return
		}

		c.JSON(http.StatusOK, envelope{Data: data, Error: nil})
	}
}
