package handler

import (
	"html/template"
	"net/http"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/notification"
)

type NotificationHandler struct {
	tmpl                *template.Template
	notificationService notification.NotificationService
}

func NewNotificationHandler(T *template.Template, NS notification.NotificationService) *NotificationHandler {
	return &NotificationHandler{tmpl: T, notificationService: NS}
}

func (ntf *NotificationHandler) GetNotification(w http.ResponseWriter, r *http.Request) {

	notf, err := ntf.notificationService.GetNotification()
	if err != nil {
		panic(err)
	}
	ntf.tmpl.ExecuteTemplate(w, "student.notification.layout", notf)

}

func (ntf *NotificationHandler) AddNotification(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		notf := entity.Notification{}

		notf.NotifyName = r.FormValue("name")
		notf.Message = r.FormValue("message")

		ntf.notificationService.AddNotification(notf)

		http.Redirect(w, r, "/teacher/register", http.StatusSeeOther)

	}

	ntf.tmpl.ExecuteTemplate(w, "teacher.notification.layout", nil)

}
