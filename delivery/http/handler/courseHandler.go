package handler

import (
	"encoding/json"
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

func (crs *CourseHandler) AdminAddCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		course := entity.Course{}
		course.CourseName = r.FormValue("coursename")
		course.CourseID, _ = strconv.Atoi(r.FormValue("courseid"))
		course.Grade, _ = strconv.Atoi(r.FormValue("grade"))

		crs.crsService.AddCourse(course)


	}

	crs.tmpl.ExecuteTemplate(w, "admin.course.new.layout", nil)



}

func (crs *CourseHandler) AdminGetCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := crs.crsService.GetCourse()
	if err != nil {
		panic(err)
	}

	crs.tmpl.ExecuteTemplate(w, "admin.course.layout", courses)

}


func (crs *CourseHandler) AdminUpdateCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("courseid")
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
		course.CourseID, _ = strconv.Atoi(r.FormValue("courseid"))
		course.CourseName = r.FormValue("coursename")
		course.Grade,_ = strconv.Atoi(r.FormValue("grade"))


		err := crs.crsService.UpdateCourse(course)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/course", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/course", http.StatusSeeOther)
	}

}

func (crs *CourseHandler) AdminDeleteCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("courseid")

		id,err := strconv.ParseInt(idRaw,0,0)

		if err != nil {
			panic(err)
		}

		 crs.crsService.DeleteCourse(int(id))


	}

	http.Redirect(w, r, "/admin/course", http.StatusSeeOther)
}


func (crs *CourseHandler) StudentGetCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := crs.crsService.GetCourse()
	if err != nil {
		panic(err)
	}

	crs.tmpl.ExecuteTemplate(w, "student.course.layout", courses)

}

func (crs *CourseHandler) FamilyGetCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := crs.crsService.GetCourse()
	if err != nil {
		panic(err)
	}

	crs.tmpl.ExecuteTemplate(w, "family.course.layout", courses)

}


func (crs *CourseHandler)AdminPostCourse(w http.ResponseWriter,r *http.Request){

	  len := r.ContentLength

	  body:= make([]byte,len)

	  r.Body.Read(body)

	  course:= entity.Course{}

	  json.Unmarshal(body,&course)


	  crs.crsService.AddCourse(course)

	   w.WriteHeader(200)

	   return
}






