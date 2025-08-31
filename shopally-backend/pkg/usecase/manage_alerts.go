package usecase

import "github.com/shopally-ai/pkg/domain"

type AlertManager struct {
	repo domain.AlertRepository
}

func NewAlertManager(repo domain.AlertRepository) *AlertManager {
	return &AlertManager{
		repo: repo,
	}
}
func (m *AlertManager) CreateAlert(alert *domain.Alert) error {
	return m.repo.CreateAlert(alert)
}

func (m *AlertManager) GetAlert(alertID string) (*domain.Alert, error) {
	return m.repo.GetAlert(alertID)
}

func (m *AlertManager) DeleteAlert(alertID string) error {
	return m.repo.DeleteAlert(alertID)
}
