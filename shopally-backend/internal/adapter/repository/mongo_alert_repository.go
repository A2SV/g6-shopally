package repository

import (
	"fmt"
	"sync"

	"github.com/shopally-ai/pkg/domain"

	"github.com/google/uuid"
)

type MockAlertRepository struct {
	alerts sync.Map // map[string]*domain.Alert
}

func NewMockAlertRepository() *MockAlertRepository {
	return &MockAlertRepository{}
}
func (r *MockAlertRepository) CreateAlert(alert *domain.Alert) error {
	alert.ID = uuid.New().String()
	r.alerts.Store(alert.ID, alert)
	return nil
}
func (r *MockAlertRepository) GetAlert(alertID string) (*domain.Alert, error) {
	if value, ok := r.alerts.Load(alertID); ok {
		if alert, ok := value.(*domain.Alert); ok {
			return alert, nil
		}
	}
	return nil, fmt.Errorf("alert with ID %s not found", alertID)
}
func (r *MockAlertRepository) DeleteAlert(alertID string) error {
	value, ok := r.alerts.Load(alertID)
	if !ok {
		return fmt.Errorf("alert with ID %s not found", alertID)
	}
	alert, ok := value.(*domain.Alert)
	if !ok {
		return fmt.Errorf("alert with ID %s not found", alertID)
	}
	alert.IsActive = false
	r.alerts.Store(alertID, alert)
	return nil
}
