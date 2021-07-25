package interfaces

import (
	"github.com/joao1559/golang-api-scaffold/models"
)

//HealthCheckUseCase ...
type HealthCheckUseCase interface {
	Check() (*models.HealthCheck, error)
}
