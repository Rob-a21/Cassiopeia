package handler

import (
	"html/template"
	"net/http"
)


type HomeHandler struct {
	tmpl        *template.Template
}

func NewHomeHandler(T *template.Template) *HomeHandler {
	return &HomeHandler{tmpl: T}
}



func (srh *LoginHandler) Home(w http.ResponseWriter, r *http.Request) {


		srh.tmpl.ExecuteTemplate(w, "mainPage.html", nil)

	}


