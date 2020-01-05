package notification

import (
	"github.com/Rob-a21/Cassiopeia/entity"
)

type NotificationRepository interface {
	GetNotification() ([]entity.Notification, error)
	AddNotification(entity.Notification) error
}
