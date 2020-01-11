package handler

import (
	"github.com/Rob-a21/Cassiopeia/models"
	"html/template"
	"net/http"
)

type LogoutHandler struct {
	tmpl         *template.Template
	logoutService models.ProfileService
}

func NewLogoutHandler(T *template.Template, PS models.ProfileService) *LogoutHandler {
	return &LogoutHandler{tmpl: T, logoutService: PS}
}

func (srh *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request) {

	  http.Redirect(w,r,"/",http.StatusSeeOther)
}