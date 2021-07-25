package usecases_test

import (
	"errors"
	"testing"

	"github.com/joao1559/golang-api-scaffold/models"
	"github.com/joao1559/golang-api-scaffold/repositories/mocks"
	"github.com/joao1559/golang-api-scaffold/usecases"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheckCheck(t *testing.T) {
	mockHealthCheckRepo := new(mocks.HealthCheckRepository)
	mockHealthCheck := models.HealthCheck{}

	t.Run("success", func(t *testing.T) {
		mockHealthCheckRepo.On("Check").Return(&mockHealthCheck, nil).Once()
		u := usecases.NewHealthCheckUseCase(mockHealthCheckRepo)

		a, err := u.Check()

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockHealthCheckRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockHealthCheckRepo.On("Check").Return(nil, errors.New("Unexpected")).Once()

		u := usecases.NewHealthCheckUseCase(mockHealthCheckRepo)

		a, err := u.Check()

		assert.Error(t, err)
		assert.Nil(t, a)

		mockHealthCheckRepo.AssertExpectations(t)
	})
}
