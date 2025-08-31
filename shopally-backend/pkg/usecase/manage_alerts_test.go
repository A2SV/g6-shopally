package usecase

import (
	"fmt"
	"github.com/shopally-ai/pkg/domain"
	"sync"
	"testing"
)

type mockAlertRepository struct {
	alerts sync.Map
}

func newMockAlertRepository() *mockAlertRepository {
	return &mockAlertRepository{}
}

func (m *mockAlertRepository) CreateAlert(alert *domain.Alert) error {
	if alert.ID == "" {
		alert.ID = "test-alert-id"
	}
	m.alerts.Store(alert.ID, alert)
	return nil
}

func (m *mockAlertRepository) GetAlert(alertID string) (*domain.Alert, error) {
	if value, ok := m.alerts.Load(alertID); ok {
		return value.(*domain.Alert), nil
	}
	return nil, fmt.Errorf("alert with ID %s not found", alertID)
}

func (m *mockAlertRepository) DeleteAlert(alertID string) error {
	if _, ok := m.alerts.Load(alertID); !ok {
		return fmt.Errorf("alert with ID %s not found", alertID)
	}
	m.alerts.Delete(alertID)
	return nil
}

func TestAlertManager_UseCases(t *testing.T) {
	mockRepo := newMockAlertRepository()
	alertManager := NewAlertManager(mockRepo)

	sampleAlert := &domain.Alert{
		UserID:      "user-123",
		ProductID:   "prod-abc",
		TargetPrice: 500.00,
	}

	var createdAlertID string
	t.Run("CreateAlert_Success", func(t *testing.T) {
		err := alertManager.CreateAlert(sampleAlert)
		if err != nil {
			t.Fatalf("CreateAlert failed: %v", err)
		}
		if sampleAlert.ID == "" {
			t.Fatal("Alert ID was not populated by the use case")
		}
		createdAlertID = sampleAlert.ID
	})
	t.Run("GetAlert_Success", func(t *testing.T) {
		retrievedAlert, err := alertManager.GetAlert(createdAlertID)
		if err != nil {
			t.Fatalf("GetAlert failed: %v", err)
		}
		if retrievedAlert.ID != createdAlertID {
			t.Errorf("Retrieved alert ID mismatch: got %s, want %s", retrievedAlert.ID, createdAlertID)
		}
	})

	t.Run("GetAlert_NotFound", func(t *testing.T) {
		_, err := alertManager.GetAlert("non-existent-id")
		if err == nil {
			t.Fatal("GetAlert for non-existent ID did not return an error")
		}
	})

	t.Run("DeleteAlert_Success", func(t *testing.T) {
		err := alertManager.DeleteAlert(createdAlertID)
		if err != nil {
			t.Fatalf("DeleteAlert failed: %v", err)
		}

		_, err = alertManager.GetAlert(createdAlertID)
		if err == nil {
			t.Fatal("Alert was not deleted as expected")
		}
	})

	t.Run("DeleteAlert_NotFound", func(t *testing.T) {
		err := alertManager.DeleteAlert("non-existent-id")
		if err == nil {
			t.Fatal("DeleteAlert for non-existent ID did not return an error")
		}
	})
}
