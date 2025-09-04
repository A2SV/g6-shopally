package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompareProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Helper to set default headers
	addDefaultHeaders := func(req *http.Request) {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept-Language", "en")
		req.Header.Set("X-Device-ID", "test-device-id")
	}

	t.Run("Success Case: Should return 200 OK with comparison data", func(t *testing.T) {
		mockUseCase := new(usecase.MockCompareProductsUseCase)
		handler := NewCompareHandler(mockUseCase)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			lang := c.GetHeader("Accept-Language")
			if lang == "" {
				lang = "en"
			}
			c.Set(string(contextkeys.RespLang), lang)
		})
		router.POST("/compare", handler.CompareProducts)

		productsToCompare := []*domain.Product{
			{ID: "ALI-123", Title: "Product A"},
			{ID: "ALI-456", Title: "Product B"},
		}
		requestBody, _ := json.Marshal(gin.H{"products": productsToCompare})
		expectedResult := map[string]interface{}{"comparison": "some comparison data"}

		mockUseCase.
			On("Execute", mock.Anything, productsToCompare).
			Return(expectedResult, nil).
			Once()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/compare", bytes.NewBuffer(requestBody))
		addDefaultHeaders(req)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &responseBody)

		assert.EqualValues(t, expectedResult, responseBody["data"])
		assert.Nil(t, responseBody["error"])
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Validation Error Case: Should return 400 Bad Request for invalid number of products", func(t *testing.T) {
		mockUseCase := new(usecase.MockCompareProductsUseCase)
		handler := NewCompareHandler(mockUseCase)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			lang := c.GetHeader("Accept-Language")
			if lang == "" {
				lang = "en"
			}
			c.Set(string(contextkeys.RespLang), lang)
		})
		router.POST("/compare", handler.CompareProducts)

		productsToCompare := []*domain.Product{
			{ID: "ALI-123", Title: "Product A"}, // Only 1 product → invalid
		}
		requestBody, _ := json.Marshal(gin.H{"products": productsToCompare})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/compare", bytes.NewBuffer(requestBody))
		addDefaultHeaders(req)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &responseBody)

		errorData := responseBody["error"].(map[string]interface{})
		assert.Equal(t, "INVALID_INPUT", errorData["code"])
		// Fixed: match the actual message returned by the handler
		assert.Contains(t, errorData["message"], "Products array must contain 2 to 4 items")

		mockUseCase.AssertNotCalled(t, "Execute")
	})

	t.Run("Malformed JSON Case: Should return 400 Bad Request for bad JSON", func(t *testing.T) {
		mockUseCase := new(usecase.MockCompareProductsUseCase)
		handler := NewCompareHandler(mockUseCase)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			lang := c.GetHeader("Accept-Language")
			if lang == "" {
				lang = "en"
			}
			c.Set(string(contextkeys.RespLang), lang)
		})
		router.POST("/compare", handler.CompareProducts)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/compare", bytes.NewBufferString("{\"invalid\":\"json\"}"))
		addDefaultHeaders(req)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockUseCase.AssertNotCalled(t, "Execute")
	})
}
