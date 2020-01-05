package handler

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/registration"
)

type RegistrationHandler struct {
	tmpl       *template.Template
	regService registration.RegistrationService
}

func NewRegistrationHandler(T *template.Template, RS registration.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{tmpl: T, regService: RS}
}

func (srh *RegistrationHandler) StudentRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		student := entity.Student{}
		student.UserName = r.FormValue("username")
		student.Password = r.FormValue("password")
		student.FirstName = r.FormValue("fname")
		student.LastName = r.FormValue("lname")
		student.ID, _ = strconv.Atoi(r.FormValue("id"))
		student.Email = r.FormValue("email")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		student.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		srh.regService.RegisterStudent(student)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/student/register", http.StatusSeeOther)

	} else {

		srh.tmpl.ExecuteTemplate(w, "student.registration.html", nil)

	}
}

func (srh *RegistrationHandler) FamilyRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		family := entity.Family{}
		family.FirstName = r.FormValue("fname")
		family.LastName = r.FormValue("lname")
		family.Password = r.FormValue("password")
		family.Phone = r.FormValue("phone")
		family.Email = r.FormValue("email")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		family.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		srh.regService.RegisterFamily(family)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "family/register", http.StatusSeeOther)

	} else {

		srh.tmpl.ExecuteTemplate(w, "family.registration.html", nil)

	}
}

func (srh *RegistrationHandler) TeacherRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		teacher := entity.Teacher{}
		teacher.UserName = r.FormValue("username")
		teacher.Password = r.FormValue("password")
		teacher.Phone = r.FormValue("phone")
		teacher.Email = r.FormValue("email")
		teacher.FirstName = r.FormValue("fname")
		teacher.LastName = r.FormValue("lname")
		teacher.TeacherID = r.FormValue("id")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		teacher.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		srh.regService.RegisterTeacher(teacher)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/teacher/register", http.StatusSeeOther)

	} else {

		srh.tmpl.ExecuteTemplate(w, "teacher.registration.html", nil)

	}
}

func (srh *RegistrationHandler) AdminRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		admin := entity.Admin{}
		admin.UserName = r.FormValue("username")
		admin.Password = r.FormValue("password")
		admin.FirstName = r.FormValue("fname")
		admin.LastName = r.FormValue("lname")
		admin.Email = r.FormValue("email")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		admin.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		srh.regService.RegisterAdmin(admin)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)

	} else {

		srh.tmpl.ExecuteTemplate(w, "admin.registration.html", nil)

	}
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "delivery", "web", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
