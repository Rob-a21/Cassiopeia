package entity

import "time"

type Notification struct{

	Message string             `json:"message""`

	NotifyName string           `json:"name""`

	NotificationDate time.Time   `json:"date""`
}