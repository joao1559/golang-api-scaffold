package interfaces

import (
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

//NotificationUseCase ...
type NotificationUseCase interface {
	GetAll(*models.Notification) ([]*models.Notification, error)
	Update(*models.Notification) error
	Read(*models.Notification) error
}

//HealthCheckUseCase ...
type HealthCheckUseCase interface {
	Check() (*models.HealthCheck, error)
}
