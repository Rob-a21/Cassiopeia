package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Rob-a21/Cassiopeia/user"
)

type ProfileHandler struct {
	tmpl           *template.Template
	profileService user.ProfileService
}

func NewProfileHandler(T *template.Template, PS user.ProfileService) *ProfileHandler {
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




func (prf *ProfileHandler) AdminGetStudent(w http.ResponseWriter, r *http.Request) {

	students, err := prf.profileService.Students()
	if err != nil {
		panic(err)
	}

	prf.tmpl.ExecuteTemplate(w, "admin.viewstudent.layout", students)



}

func (prf *ProfileHandler) AdminDeleteStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = prf.profileService.DeleteStudent(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/student", http.StatusSeeOther)
}




func (prf *ProfileHandler) AdminGetTeacher(w http.ResponseWriter, r *http.Request) {

	teachers, err := prf.profileService.Teachers()
	if err != nil {
		panic(err)
	}

	prf.tmpl.ExecuteTemplate(w, "admin.viewteacher.layout", teachers)



}

func (prf *ProfileHandler) AdminDeleteTeacher(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		id := r.URL.Query().Get("id")

		//id, err := strconv.Atoi(idRaw)

		//if err != nil {
		//	panic(err)
		//}

		 prf.profileService.DeleteTeacher(id)



	}

	http.Redirect(w, r, "/admin/teacher", http.StatusSeeOther)
}
