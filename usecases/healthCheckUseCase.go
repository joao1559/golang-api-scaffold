package usecases

import (
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/interfaces"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

type healthcheckUseCase struct {
	healthcheckRepo interfaces.HealthCheckRepository
}

//NewHealthCheckUseCase will create new an healthcheckUsecase object representation of usecase.HealthCheckUsecase interface
func NewHealthCheckUseCase(h interfaces.HealthCheckRepository) interfaces.HealthCheckUseCase {
	return &healthcheckUseCase{
		healthcheckRepo: h,
	}
}

func (h *healthcheckUseCase) Check() (*models.HealthCheck, error) {
	res, err := h.healthcheckRepo.Check()
	if err != nil {
		return nil, err
	}
	return res, nil
}
