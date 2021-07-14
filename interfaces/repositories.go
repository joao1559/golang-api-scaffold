package interfaces

import (
	"github.com/joao1559/golang-api-scaffold/models"
)

//HealthCheckRepository ...
type HealthCheckRepository interface {
	Check() (*models.HealthCheck, error)
}
