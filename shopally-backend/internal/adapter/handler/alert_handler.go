// Package handler provides HTTP handlers for alert-related endpoints.
package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/usecase"
)

// AlertHandler handles HTTP requests for alert operations.
type AlertHandler struct {
	alertManager *usecase.AlertManager
}

// NewAlertHandler creates a new AlertHandler with the given AlertManager.
func NewAlertHandler(am *usecase.AlertManager) *AlertHandler {
	return &AlertHandler{
		alertManager: am,
	}
}

// successResponse represents a standard API response structure.
type successResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

// createAlertPayload represents the expected payload for creating an alert.
type createAlertPayload struct {
	ProductID    string  `json:"productId"`
	DeviceID     string  `json:"deviceId"`
	CurrentPrice float64 `json:"currentPrice"`
}

// CreateAlertHandler handles POST requests to create a new alert.
func (h *AlertHandler) CreateAlertHandler(c *gin.Context) {
	var payload createAlertPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newAlert := &domain.Alert{
		ProductID:    payload.ProductID,
		DeviceID:     payload.DeviceID,
		CurrentPrice: payload.CurrentPrice,
		IsActive:     true,
	}

	if err := h.alertManager.CreateAlert(newAlert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create alert: %v", err)})
		return
	}

	response := successResponse{
		Data: map[string]string{
			"status":  "Alert created successfully",
			"alertId": newAlert.ID,
		},
		Error: nil,
	}
	c.JSON(http.StatusCreated, response)
}

// GetAlertHandler handles GET requests to retrieve an alert by its ID.
func (h *AlertHandler) GetAlertHandler(c *gin.Context) {
	alertID := c.Param("id")
	alert, err := h.alertManager.GetAlert(alertID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Failed to retrieve alert: %v", err)})
		return
	}

	response := successResponse{
		Data:  alert,
		Error: nil,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteAlertHandler handles DELETE requests to remove an alert by its ID.
func (h *AlertHandler) DeleteAlertHandler(c *gin.Context) {
	alertID := c.Param("id")
	if err := h.alertManager.DeleteAlert(alertID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Failed to delete alert: %v", err)})
		return
	}

	response := successResponse{
		Data: map[string]string{
			"status": "Alert deleted successfully",
		},
		Error: nil,
	}
	c.JSON(http.StatusOK, response)
}
