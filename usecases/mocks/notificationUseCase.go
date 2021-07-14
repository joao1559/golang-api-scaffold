package mocks

import (
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
	"github.com/stretchr/testify/mock"
)

//NotificationUseCase ...
type NotificationUseCase struct {
	mock.Mock
}

//GetAll ...
func (u *NotificationUseCase) GetAll(c *models.Notification) ([]*models.Notification, error) {
	ret := u.Called(c)
	var r0 []*models.Notification
	if rf, ok := ret.Get(0).(func(*models.Notification) []*models.Notification); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Notification) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//Update ...
func (u *NotificationUseCase) Update(c *models.Notification) error {
	ret := u.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Notification) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//Read ...
func (u *NotificationUseCase) Read(c *models.Notification) error {
	ret := u.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Notification) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

