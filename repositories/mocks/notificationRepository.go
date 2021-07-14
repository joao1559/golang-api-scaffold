package mocks

import (
	"github.com/stretchr/testify/mock"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

//NotificationRepository ...
type NotificationRepository struct {
	mock.Mock
}

//GetAll ...
func (r *NotificationRepository) GetAll(not *models.Notification) ([]*models.Notification, error) {
	ret := r.Called(not)
	var r0 []*models.Notification
	if rf, ok := ret.Get(0).(func(*models.Notification) []*models.Notification); ok {
		r0 = rf(not)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Notification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Notification) error); ok {
		r1 = rf(not)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//Update ...
func (r *NotificationRepository) Update(m *models.Notification) error {
	ret := r.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Notification) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//Read ...
func (r *NotificationRepository) Read(m *models.Notification) error {
	ret := r.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Notification) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}