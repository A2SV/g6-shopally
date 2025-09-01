package repository

import (
	"testing"

	"github.com/shopally-ai/internal/mocks"
	"github.com/shopally-ai/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AlertRepositorySuite struct {
	suite.Suite
	mockRepo    *mocks.AlertRepository
	sampleAlert *domain.Alert
}

func (s *AlertRepositorySuite) SetupTest() {
	s.mockRepo = new(mocks.AlertRepository)
	s.sampleAlert = &domain.Alert{
		DeviceID:     "device-123",
		ProductID:    "product-abc",
		CurrentPrice: 500.00,
		IsActive:     true,
	}
}

func (s *AlertRepositorySuite) TestCreateAlert_Success() {
	s.mockRepo.On("CreateAlert", s.sampleAlert).Return(nil)
	err := s.mockRepo.CreateAlert(s.sampleAlert)
	s.NoError(err)
	s.mockRepo.AssertCalled(s.T(), "CreateAlert", s.sampleAlert)
}

func (s *AlertRepositorySuite) TestGetAlert_Success() {
	s.mockRepo.On("GetAlert", "test-alert-id").Return(s.sampleAlert, nil)
	alert, err := s.mockRepo.GetAlert("test-alert-id")
	s.NoError(err)
	s.Equal(s.sampleAlert, alert)
	s.mockRepo.AssertCalled(s.T(), "GetAlert", "test-alert-id")
}

func (s *AlertRepositorySuite) TestGetAlert_NotFound() {
	s.mockRepo.On("GetAlert", "non-existent-id").Return(nil, assert.AnError)
	alert, err := s.mockRepo.GetAlert("non-existent-id")
	s.Error(err)
	s.Nil(alert)
	s.mockRepo.AssertCalled(s.T(), "GetAlert", "non-existent-id")
}

func (s *AlertRepositorySuite) TestDeleteAlert_Success() {
	s.mockRepo.On("DeleteAlert", "test-alert-id").Return(nil)
	err := s.mockRepo.DeleteAlert("test-alert-id")
	s.NoError(err)
	s.mockRepo.AssertCalled(s.T(), "DeleteAlert", "test-alert-id")
}

func (s *AlertRepositorySuite) TestDeleteAlert_NotFound() {
	s.mockRepo.On("DeleteAlert", "non-existent-id").Return(assert.AnError)
	err := s.mockRepo.DeleteAlert("non-existent-id")
	s.Error(err)
	s.mockRepo.AssertCalled(s.T(), "DeleteAlert", "non-existent-id")
}

func TestAlertRepositorySuite(t *testing.T) {
	suite.Run(t, new(AlertRepositorySuite))
}

func TestMockAlertRepository(t *testing.T) {
	repo := NewMockAlertRepository()

	sampleAlert := &domain.Alert{
		DeviceID:     "device-123",
		ProductID:    "product-abc",
		CurrentPrice: 500.00,
		IsActive:     true,
	}

	var createdAlertID string

	t.Run("CreateAlert_Success", func(t *testing.T) {
		err := repo.CreateAlert(sampleAlert)
		if err != nil {
			t.Fatalf("CreateAlert failed with error: %v", err)
		}

		if sampleAlert.ID == "" {
			t.Fatal("CreateAlert did not assign an ID to the alert")
		}
		createdAlertID = sampleAlert.ID
	})

	t.Run("GetAlert_Success", func(t *testing.T) {
		retrievedAlert, err := repo.GetAlert(createdAlertID)
		if err != nil {
			t.Fatalf("GetAlert failed with error: %v", err)
		}

		if retrievedAlert.ID != createdAlertID {
			t.Errorf("Retrieved alert ID mismatch: got %s, want %s", retrievedAlert.ID, createdAlertID)
		}

		if retrievedAlert.DeviceID != sampleAlert.DeviceID {
			t.Errorf("Retrieved alert DeviceID mismatch: got %s, want %s", retrievedAlert.DeviceID, sampleAlert.DeviceID)
		}
	})

	t.Run("GetAlert_NotFound", func(t *testing.T) {
		_, err := repo.GetAlert("non-existent-id")
		if err == nil {
			t.Fatal("GetAlert for a non-existent ID did not return an error")
		}
	})

	t.Run("DeleteAlert_Success", func(t *testing.T) {
		err := repo.DeleteAlert(createdAlertID)
		if err != nil {
			t.Fatalf("DeleteAlert failed with error: %v", err)
		}

		alert, err := repo.GetAlert(createdAlertID)
		if err != nil {
			t.Fatalf("GetAlert failed after delete: %v", err)
		}
		if alert.IsActive {
			t.Errorf("expected IsActive to be false after deletion, got true")
		}
	})

	t.Run("DeleteAlert_NotFound", func(t *testing.T) {
		err := repo.DeleteAlert("non-existent-id")
		if err == nil {
			t.Fatal("DeleteAlert for a non-existent ID did not return an error")
		}
	})
}
