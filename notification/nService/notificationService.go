package nService

import (
	"github.com/robi_a21/Cassiopeia/entity"
	"github.com/robi_a21/Cassiopeia/notification"
)

type NotificationServiceImpl struct {
	notificationRepository notification.NotificationRepository
}

func NewNotificationServiceImpl(nrpo notification.NotificationRepository) *NotificationServiceImpl {
	return &NotificationServiceImpl{notificationRepository: nrpo}
}

func (ntf *NotificationServiceImpl) AddNotification(notf entity.Notification) error {

	err := ntf.notificationRepository.AddNotification(notf)

	if err != nil {
		return err
	}

	return nil
}


func (ss *NotificationServiceImpl) GetNotification() ([]entity.Notification, error) {

	notn, err := ss.notificationRepository.GetNotification()

	if err != nil {
		return nil, err
	}

	return notn, nil
}





