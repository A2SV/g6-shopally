package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/internal/adapter/repository"
	"github.com/shopally-ai/pkg/usecase"
)

func TestAlertHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := repository.NewMockAlertRepository()
	alertManager := usecase.NewAlertManager(mockRepo)
	alertHandler := NewAlertHandler(alertManager)

	var alertID string

	t.Run("CreateAlertHandler", func(t *testing.T) {
		payload := []byte(`{"deviceId": "device-123", "productId": "prod-abc", "currentPrice": 500.00}`)
		req := httptest.NewRequest("POST", "/alerts", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = req

		alertHandler.CreateAlertHandler(c)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var res successResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}

		dataMap, ok := res.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("response data is not a map: got %T", res.Data)
		}
		if status, ok := dataMap["status"].(string); !ok || status != "Alert created successfully" {
			t.Errorf("unexpected status message: got %v", status)
		}
		id, ok := dataMap["alertId"].(string)
		if !ok || id == "" {
			t.Errorf("missing or invalid alertId in response: got %v", id)
		} else {
			alertID = id
		}
	})

	t.Run("GetAlertHandler", func(t *testing.T) {
		if alertID == "" {
			t.Fatal("alertID was not set in previous test")
		}
		req := httptest.NewRequest("GET", "/alerts/"+alertID, nil)
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Params = gin.Params{{Key: "id", Value: alertID}}

		alertHandler.GetAlertHandler(c)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var res successResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}

		alertData, ok := res.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("response data is not a map: got %T", res.Data)
		}
		if alertId, ok := alertData["alertId"].(string); !ok || alertId != alertID {
			t.Errorf("unexpected alertId in response: got %v", alertData["alertId"])
		}
	})

	t.Run("DeleteAlertHandler", func(t *testing.T) {
		if alertID == "" {
			t.Fatal("alertID was not set in previous test")
		}
		req := httptest.NewRequest("DELETE", "/alerts/"+alertID, nil)
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Params = gin.Params{{Key: "id", Value: alertID}}

		alertHandler.DeleteAlertHandler(c)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var res successResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}

		dataMap, ok := res.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("response data is not a map: got %T", res.Data)
		}
		if status, ok := dataMap["status"].(string); !ok || status != "Alert deleted successfully" {
			t.Errorf("unexpected status message: got %v", status)
		}
	})

	t.Run("GetDeletedAlertIsInactive", func(t *testing.T) {
		if alertID == "" {
			t.Fatal("alertID was not set in previous test")
		}
		req := httptest.NewRequest("GET", "/alerts/"+alertID, nil)
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Params = gin.Params{{Key: "id", Value: alertID}}

		alertHandler.GetAlertHandler(c)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var res successResponse
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("could not decode response body: %v", err)
		}

		alertData, ok := res.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("response data is not a map: got %T", res.Data)
		}
		isActive, ok := alertData["isActive"].(bool)
		if !ok {
			t.Fatalf("isActive field missing or not a bool: got %v", alertData["isActive"])
		}
		if isActive {
			t.Errorf("expected isActive to be false after deletion, got true")
		}
	})
}
