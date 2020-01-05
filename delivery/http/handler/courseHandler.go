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
