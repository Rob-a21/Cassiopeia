package handler

import (
	"github.com/Rob-a21/Cassiopeia/models"
	"html/template"
	"net/http"
)


type HomeHandler struct {
	tmpl         *template.Template
	HomeService models.ProfileService
}

func NewHomeHandler(T *template.Template, PS models.ProfileService) *HomeHandler {
	return &HomeHandler{tmpl: T, HomeService: PS}
}

func (srh *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {


	srh.tmpl.ExecuteTemplate(w,"mainpage.layout",nil)
}