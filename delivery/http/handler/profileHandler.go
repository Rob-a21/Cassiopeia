package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Rob-a21/Cassiopeia/models"
)

type ProfileHandler struct {
	tmpl           *template.Template
	profileService models.ProfileService
}

func NewProfileHandler(T *template.Template, NS models.ProfileService) *ProfileHandler {
	return &ProfileHandler{tmpl: T, profileService: NS}
}

func (prf *ProfileHandler) StudentsProfile(w http.ResponseWriter, r *http.Request) {

	students, err := prf.profileService.Students()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "student.profile.layout", students)

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

		prf.tmpl.ExecuteTemplate(w, "student.profile.layout", student)

	}

}


func (prf *ProfileHandler) TeachersProfile(w http.ResponseWriter, r *http.Request) {

	teacher, err := prf.profileService.Teachers()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "teacher.profile.layout", teacher)

}


func (prf *ProfileHandler) TeacherProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {

			panic(err)
		}

		teacher, err := prf.profileService.Teacher(id)

		if err != nil {
			panic(err)
		}

		prf.tmpl.ExecuteTemplate(w, "teacher.profile.layout", teacher)

	}
}
func (prf *ProfileHandler) AdminProfile(w http.ResponseWriter, r *http.Request) {

	admin, err := prf.profileService.Admins()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "admin.profile.layout", admin)

}

func (prf *ProfileHandler) FamiliesProfile(w http.ResponseWriter, r *http.Request) {

	family, err := prf.profileService.Families()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "family.profile.layout", family)

}
func (prf *ProfileHandler) FamilyProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {

			panic(err)
		}

		student, err := prf.profileService.Family(id)

		if err != nil {
			panic(err)
		}

		prf.tmpl.ExecuteTemplate(w, "family.profile.html", student)

	}
}
func (prf *ProfileHandler) EmailExists(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		email := r.URL.Query().Get("id")

		prf.profileService.EmailExists(email)

	}

	http.Redirect(w, r, "/admin/student", http.StatusSeeOther)
}


func (prf *ProfileHandler) AdminGetStudent(w http.ResponseWriter, r *http.Request) {

	students, err := prf.profileService.Students()
	if err != nil {
		panic(err)
	}

	prf.tmpl.ExecuteTemplate(w, "admin.view.student.layout", students)

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

	prf.tmpl.ExecuteTemplate(w, "admin.view.teacher.layout", teachers)

}

func (prf *ProfileHandler) AdminDeleteTeacher(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		prf.profileService.DeleteTeacher(id)

	}

	http.Redirect(w, r, "/admin/teacher", http.StatusSeeOther)
}
