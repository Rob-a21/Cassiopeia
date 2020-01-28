package handler

import (
	"encoding/json"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
	"html/template"
	"net/http"
	"strconv"
)

type AssessmentHandler struct {
	tmpl       *template.Template
	assService models.AssessmentService
}

func NewAssessmentHandler(T *template.Template, AS models.AssessmentService) *AssessmentHandler {
	return &AssessmentHandler{tmpl: T, assService: AS}
}


func (as *AssessmentHandler) AssessmentsOfOneGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		grade := r.FormValue("grade")
		assess, err := as.assService.Assessments(grade)

		if err != nil {
			panic(err)
		}

		_ = as.tmpl.ExecuteTemplate(w, "teacher.assessment.layout", assess)

	}

}




func (as *AssessmentHandler) UpdateGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		assessment := entity.Assessment{}
		assessment.Value,_ = strconv.Atoi(r.FormValue("value"))
		assessment.SubjectID,_ = strconv.Atoi(r.FormValue("subjectid"))
		assessment.StudentID,_ =  strconv.Atoi(r.FormValue("studentid"))
		assessment.Grade,_ = strconv.Atoi(r.FormValue("grade"))


		err := as.assService.UpdateGrade(assessment)

		if err != nil {
			panic(err)
		}


	}
	_ = as.tmpl.ExecuteTemplate(w, "admin.grade.update.layout", nil)


}

func (as *AssessmentHandler) DeleteGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		SubjectID,_ := strconv.Atoi(r.FormValue("subjectid"))
		StudentID,_ := strconv.Atoi(r.FormValue("studentid"))

		err := as.assService.DeleteGrade(StudentID,SubjectID)

		if err != nil {
			panic(err)
		}


	}
	_ = as.tmpl.ExecuteTemplate(w, "admin.grade.layout", nil)


}

func (as *AssessmentHandler) DeleteGrades(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		StudentID,_ := strconv.Atoi(r.FormValue("studentid"))

		err := as.assService.DeleteGrades(StudentID)

		if err != nil {
			panic(err)
		}


	}
	_ = as.tmpl.ExecuteTemplate(w, "admin.grade.layout", nil)


}


func (as *AssessmentHandler) StoreGrade(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		assessment := entity.Assessment{}
		assessment.Value,_ = strconv.Atoi(r.FormValue("value"))
		assessment.SubjectID,_= strconv.Atoi(r.FormValue("subjectid"))
		assessment.StudentID,_ = strconv.Atoi(r.FormValue("studentid"))
		assessment.Grade,_ = strconv.Atoi(r.FormValue("grade"))


		err := as.assService.StoreGrade(assessment)

		if err != nil {
			panic(err)
		}


	}
	_ = as.tmpl.ExecuteTemplate(w, "teacher.assessment.new.layout", nil)


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


func (gr *AssessmentHandler)ApiTeacherPostGrade(w http.ResponseWriter,r *http.Request){

	len := r.ContentLength

	body:= make([]byte,len)

	_, _ = r.Body.Read(body)

	assessment:= entity.Assessment{}

	_ = json.Unmarshal(body, &assessment)

	_ = gr.assService.StoreGrade(assessment)

	w.WriteHeader(200)

	return
}
