package repository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
)

type PsqlNotificationRepositoryImpl struct {
	conn *sql.DB
}

func NewPsqlNotificationRepositoryImpl(Conn *sql.DB) *PsqlNotificationRepositoryImpl {
	return &PsqlNotificationRepositoryImpl{conn: Conn}
}

func (ntf *PsqlNotificationRepositoryImpl) AddNotification(notf entity.Notification) error {

	_, err := ntf.conn.Exec("insert into notification (notifyname,message,notifieddate) values($1, $2,$3)", notf.NotifyName, notf.Message, notf.NotificationDate)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

func (ntf *PsqlNotificationRepositoryImpl) GetNotification() ([]entity.Notification, error) {

	rows, err := ntf.conn.Query("SELECT * FROM notification;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	notifications := []entity.Notification{}

	for rows.Next() {
		notification := entity.Notification{}
		err = rows.Scan(&notification.NotifyName, &notification.Message, &notification.NotificationDate)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, err
}

func (ntf *PsqlNotificationRepositoryImpl) Notification(id string) (entity.Notification, error) {

	row := ntf.conn.QueryRow("SELECT * FROM notification WHERE notifyname = $1", id)

	notification := entity.Notification{}

	err := row.Scan(&notification.Message, &notification.NotificationDate)
	if err != nil {
		return notification,err
	}

	return notification, nil
}
