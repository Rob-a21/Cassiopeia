package handler

import (
	"github.com/Rob-a21/Cassiopeia/assessment"
	"github.com/Rob-a21/Cassiopeia/entity"
	"html/template"
	"net/http"
	"strconv"
)

type AssessmentHandler struct {
	tmpl       *template.Template
	assService assessment.AssessmentService
}

func NewAssessmentHandler(T *template.Template, AS assessment.AssessmentService) *AssessmentHandler {
	return &AssessmentHandler{tmpl: T, assService: AS}
}


func (as *AssessmentHandler) AssessmentsOfOneGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		grade := r.FormValue("grade")
		assess, err := as.assService.Assessments(grade)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", assess)

	}

}

func (as *AssessmentHandler) SingleStudentAssessments(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		id,_ := strconv.Atoi(r.FormValue("studentid"))
		assess, err := as.assService.SingleStudentAssessment(id)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", assess)

	}

}

func (as *AssessmentHandler) SingleAssessment(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		assessment := entity.Assessment{}
		assessment.Value,_ = strconv.Atoi(r.FormValue("value"))
		assessment.Grade = r.FormValue("grade")
		assessment.SubjectID = r.FormValue("subjectid")
		assessment.StudentID = r.FormValue("studentid")

		assess, err := as.assService.Assessment(assessment)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", assess)

	}

}

func (as *AssessmentHandler) UpdateGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		assessment := entity.Assessment{}
		assessment.Value,_ = strconv.Atoi(r.FormValue("value"))
		assessment.Grade = r.FormValue("grade")
		assessment.SubjectID = r.FormValue("subjectid")
		assessment.StudentID = r.FormValue("studentid")

		err := as.assService.UpdateGrade(assessment)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", nil)

	}

}

func (as *AssessmentHandler) DeleteGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		SubjectID,_ := strconv.Atoi(r.FormValue("subjectid"))
		StudentID,_ := strconv.Atoi(r.FormValue("studentid"))

		err := as.assService.DeleteGrade(StudentID,SubjectID)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", nil)

	}

}

func (as *AssessmentHandler) DeleteGrades(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		StudentID,_ := strconv.Atoi(r.FormValue("studentid"))

		err := as.assService.DeleteGrades(StudentID)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", nil)

	}

}


func (as *AssessmentHandler) StoreGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		assessment := entity.Assessment{}
		assessment.Value,_ = strconv.Atoi(r.FormValue("value"))
		assessment.Grade = r.FormValue("grade")
		assessment.SubjectID = r.FormValue("subjectid")
		assessment.StudentID = r.FormValue("studentid")

		err := as.assService.StoreGrade(assessment)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", nil)

	}

}

func (as *AssessmentHandler) IsQualified(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		StudentID,_ := strconv.Atoi(r.FormValue("studentid"))

		val, err := as.assService.IsQualified(StudentID)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "admin.course.new.layout", val)

	}

}
