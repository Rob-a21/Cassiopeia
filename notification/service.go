package notification

import (
	"github.com/Rob-a21/Cassiopeia/entity"
)

type NotificationService interface {
	GetNotification() ([]entity.Notification, error)
	AddNotification(entity.Notification) error
}
