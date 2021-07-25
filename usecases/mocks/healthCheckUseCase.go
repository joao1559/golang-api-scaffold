package mocks

import (
	"github.com/joao1559/golang-api-scaffold/models"
	"github.com/stretchr/testify/mock"
)

//HealthCheckUseCase is an autogenerated mock type for the usecase type
type HealthCheckUseCase struct {
	mock.Mock
}

//Check provides a mock function with given fields:
func (_m *HealthCheckUseCase) Check() (*models.HealthCheck, error) {
	ret := _m.Called()

	var r0 *models.HealthCheck
	if rf, ok := ret.Get(0).(func() *models.HealthCheck); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.HealthCheck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
