package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/repositories/mocks"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/usecases"

	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

func TestGetAll(t *testing.T) {
	mockNotificationRepo := new(mocks.NotificationRepository)
	mockNotification := &models.Notification{
		ID:               1,
		Titulo:           "Notificacao",
		Mensagem:         "Nova notificacao",
		Link:             "http://localhost:8000/",
		Usuario:          "rd_moge",
		DataEnvio:        "",
		DataVisualizacao: "",
		DataCadastro:     "2019-10-01",
		UsuarioCadastro:  "rd_moge",
		DataAlteracao:    "",
		UsuarioAlteracao: "",
	}
	mockList := make([]*models.Notification, 0)
	mockList = append(mockList, mockNotification)

	t.Run("success", func(t *testing.T) {
		mockNotificationRepo.On("GetAll", mockNotification).Return(mockList, nil).Once()
		u := usecases.NewNotificationUseCase(mockNotificationRepo)
		list, err := u.GetAll(mockNotification)
		assert.NotEmpty(t, list)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockList))
		mockNotificationRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockNotificationRepo.On("GetAll", mockNotification).Return(nil, models.ErrNotFound).Once()
		u := usecases.NewNotificationUseCase(mockNotificationRepo)
		list, err := u.GetAll(mockNotification)
		assert.Empty(t, list)
		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockNotificationRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockNotificationRepo := new(mocks.NotificationRepository)
	mockNotification := &models.Notification{
		ID:               1,
		Titulo:           "Notificacao",
		Mensagem:         "Nova notificacao",
		Link:             "http://localhost:8000/",
		Usuario:          "rd_moge",
		DataEnvio:        "",
		DataVisualizacao: "",
		DataCadastro:     "2019-10-01",
		UsuarioCadastro:  "rd_moge",
		DataAlteracao:    "",
		UsuarioAlteracao: "",
	}

	t.Run("success", func(t *testing.T) {
		mockNotificationRepo.On("Update", mockNotification).Return(nil).Once()

		u := usecases.NewNotificationUseCase(mockNotificationRepo)

		err := u.Update(mockNotification)

		assert.NoError(t, err)
		mockNotificationRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockNotificationRepo.On("Update", mockNotification).Return(models.ErrNotFound)

		u := usecases.NewNotificationUseCase(mockNotificationRepo)

		err := u.Update(mockNotification)

		assert.Error(t, err)
		mockNotificationRepo.AssertExpectations(t)
	})
}

func TestRead(t *testing.T) {
	mockNotificationRepo := new(mocks.NotificationRepository)
	mockNotification := &models.Notification{
		ID:               1,
		Titulo:           "Notificacao",
		Mensagem:         "Nova notificacao",
		Link:             "http://localhost:8000/",
		Usuario:          "rd_moge",
		DataEnvio:        "",
		DataVisualizacao: "",
		DataCadastro:     "2019-10-01",
		UsuarioCadastro:  "rd_moge",
		DataAlteracao:    "",
		UsuarioAlteracao: "",
	}

	t.Run("success", func(t *testing.T) {
		mockNotificationRepo.On("Read", mockNotification).Return(nil).Once()

		u := usecases.NewNotificationUseCase(mockNotificationRepo)

		err := u.Read(mockNotification)

		assert.NoError(t, err)
		mockNotificationRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockNotificationRepo.On("Read", mockNotification).Return(models.ErrNotFound)

		u := usecases.NewNotificationUseCase(mockNotificationRepo)

		err := u.Read(mockNotification)

		assert.Error(t, err)
		mockNotificationRepo.AssertExpectations(t)
	})
}
