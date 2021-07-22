package interfaces

import (
	"github.com/joao1559/golang-api-scaffold/models"
)

//HealthCheckUseCase ...
type HealthCheckUseCase interface {
	Check() (*models.HealthCheck, error)
}

//NotesUseCase ...
type NoteUseCase interface {
	Get() ([]*models.Note, error)
	Insert([]byte) error
}
