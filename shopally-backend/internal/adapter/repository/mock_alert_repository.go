package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/shopally-ai/pkg/domain"
)

// MockAlertRepository is a simple in-memory implementation used by unit tests.
type MockAlertRepository struct {
	alerts sync.Map // key: ID, value: *domain.Alert
}

func NewMockAlertRepository() *MockAlertRepository {
	return &MockAlertRepository{}
}

func (r *MockAlertRepository) CreateAlert(alert *domain.Alert) error {
	if alert.ID == "" {
		alert.ID = uuid.New().String()
	}
	if !alert.IsActive {
		alert.IsActive = true
	}
	// store a copy to avoid external mutation side effects
	copy := *alert
	r.alerts.Store(alert.ID, &copy)
	return nil
}

func (r *MockAlertRepository) GetAlert(alertID string) (*domain.Alert, error) {
	if v, ok := r.alerts.Load(alertID); ok {
		// return a copy to avoid callers mutating internal state
		a := *(v.(*domain.Alert))
		return &a, nil
	}
	return nil, errors.New("alert not found")
}

func (r *MockAlertRepository) DeleteAlert(alertID string) error {
	if v, ok := r.alerts.Load(alertID); ok {
		a := v.(*domain.Alert)
		a.IsActive = false
		r.alerts.Store(alertID, a)
		return nil
	}
	return errors.New("alert not found")
}
