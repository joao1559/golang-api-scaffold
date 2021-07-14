package interfaces

import (
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

//NotificationRepository ...
type NotificationRepository interface {
	GetAll(*models.Notification) ([]*models.Notification, error)
	Update(*models.Notification) error
	Read(*models.Notification) error
}

//HealthCheckRepository ...
type HealthCheckRepository interface {
	Check() (*models.HealthCheck, error)
}
