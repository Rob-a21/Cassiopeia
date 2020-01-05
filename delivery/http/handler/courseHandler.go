package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Rob-a21/Cassiopeia/course"
	"github.com/Rob-a21/Cassiopeia/entity"
)

type CourseHandler struct {
	tmpl       *template.Template
	crsService course.CourseService
}

func NewCourseHandler(T *template.Template, CS course.CourseService) *CourseHandler {
	return &CourseHandler{tmpl: T, crsService: CS}
}

func (crs *CourseHandler) CourseAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		course := entity.Course{}
		course.CourseName = r.FormValue("coursename")
		course.CourseID, _ = strconv.Atoi(r.FormValue("courseid"))
		course.Grade, _ = strconv.Atoi(r.FormValue("grade"))

		crs.crsService.AddCourse(course)

	}

	crs.tmpl.ExecuteTemplate(w, "admin.course.layout", nil)

}

func (crs *CourseHandler) GetCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := crs.crsService.GetCourse()
	if err != nil {
		panic(err)
	}

	crs.tmpl.ExecuteTemplate(w, "student.course.layout", courses)

}

func (crs *CourseHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		courses, err := crs.crsService.Course(id)

		if err != nil {
			panic(err)
		}

		crs.tmpl.ExecuteTemplate(w, "admin.course.update.layout", courses)

	} else if r.Method == http.MethodPost {

		course := entity.Course{}
		course.ID, _ = strconv.Atoi(r.FormValue("courseid"))
		course.Name = r.FormValue("coursename")

		err = crs.crsService.UpdateCourse(course)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/course", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/course", http.StatusSeeOther)
	}

}

func (crs *CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		crs.crsService.DeleteCourse(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/course", http.StatusSeeOther)
}
