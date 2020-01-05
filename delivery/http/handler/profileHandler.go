package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Rob-a21/Cassiopeia/profile"
)

type ProfileHandler struct {
	tmpl           *template.Template
	profileService profile.ProfileService
}

func NewProfileHandler(T *template.Template, PS profile.ProfileService) *ProfileHandler {
	return &ProfileHandler{tmpl: T, profileService: PS}
}

func (prf *ProfileHandler) StudentsProfile(w http.ResponseWriter, r *http.Request) {

	students, err := prf.profileService.Students()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "student.index.layout", students)

}

func (prf *ProfileHandler) StudentProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {

			panic(err)
		}

		student, err := prf.profileService.Student(id)

		if err != nil {
			panic(err)
		}

		prf.tmpl.ExecuteTemplate(w, "student.index.html", student)

	}

}
