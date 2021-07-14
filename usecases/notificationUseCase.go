package usecases

import (
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/interfaces"
	"gitlab.luizalabs.com/luizalabs/tio-patinhas-notificacao-api/models"
)

type notificationUseCase struct {
	notificationRepository interfaces.NotificationRepository
}

//NewNotificationUseCase ...
func NewNotificationUseCase(h interfaces.NotificationRepository) interfaces.NotificationUseCase {
	return &notificationUseCase{
		notificationRepository: h,
	}
}

func (rel *notificationUseCase) GetAll(not *models.Notification) ([]*models.Notification, error) {
	res, err := rel.notificationRepository.GetAll(not)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rel *notificationUseCase) Update(not *models.Notification) error {
	err := rel.notificationRepository.Update(not)
	if err != nil {
		return err
	}

	return nil
}

func (rel *notificationUseCase) Read(not *models.Notification) error {
	err := rel.notificationRepository.Read(not)
	if err != nil {
		return err
	}

	return nil
}