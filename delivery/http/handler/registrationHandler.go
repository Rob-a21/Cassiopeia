package handler

import(
   "html/template"
	"net/http"
	"strconv"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/registration"

)


type StudentRegistrationHandler struct {
	tmpl        *template.Template
	regService registration.StudentService
}

func NewStudentRegistrationHandler(T *template.Template, RS registration.StudentService) *StudentRegistrationHandler {
	return &StudentRegistrationHandler{tmpl: T, regService: RS}
}


func (srh *StudentRegistrationHandler) RegistrationNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password :=	r.FormValue("password")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		id,err:= strconv.Atoi(r.FormValue("id"))
		image := r.FormValue("image")
		gender := r.FormValue("gender")
		grade := r.FormValue("grade")
		phone := r.FormValue("phone")
		email := r.FormValue("email")
		job := r.FormValue("type")


		if job == "student"{
			student := &entity.Student{}

			student.UserName = username
			student.Password = password
			student.FirstName = firstname
			student.LastName = lastname
			student.ID,_ = id,err
			student.Image = image
			student.Gender = gender
			student.Grade = grade
			student.Phone =  phone
			student.Email = email

			http.Redirect(w, r, "/student", http.StatusSeeOther)
		} else if job =="Teacher"{
			teacher := &entity.Teacher{}

			teacher.UserName = username
			teacher.Password = password
			teacher.FirstName = firstname
			teacher.LastName = lastname
			teacher.TeacherID,_ = id,err
			teacher.Image = image
			teacher.Gender = gender
			teacher.Grade = grade
			teacher.Phone =  phone
			teacher.Email = email

			http.Redirect(w, r, "/teacher", http.StatusSeeOther)
		}


	} else {

		_ = srh.tmpl.ExecuteTemplate(w, "signup.html", nil)

	}


}

