package interfaces

import (
	"github.com/joao1559/golang-api-scaffold/models"
)

//HealthCheckRepository ...
type HealthCheckRepository interface {
	Check() (*models.HealthCheck, error)
}

//NoteRepository ...
type NoteRepository interface {
	Get() ([]*models.Note, error)
	Insert(note models.Note) error
}
