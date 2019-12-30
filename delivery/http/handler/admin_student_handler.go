package handler

import (
	"encoding/json"
	"github.com/Rob-a21/Cassiopeia/entity"
	"net/http"

	"github.com/Rob-a21/Cassiopeia/registration"
	"github.com/julienschmidt/httprouter"
)

// AdminStudentHandler is used to implement role related http requests
type AdminStudentHandler struct {
	studentService registration.StudentService
}

// NewAdminStudentHandler initializes and returns new StudentRoleHandler object
func NewAdminStudentHandler(regSrv registration.StudentService) *AdminStudentHandler {
	return &AdminStudentHandler{studentService: regSrv}
}

// GetRoles handles GET /v1/admin/roles requests
func (arh *AdminStudentHandler) GetStudents(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	roles, errs := arh.studentService.Students()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleRole handles GET /v1/admin/roles/:id requests
func (arh *AdminStudentHandler) GetSingleStudent(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	students, errs := arh.studentService.Student(1)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(students, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// PutRole handles PUT /v1/admin/roles/:id requests
func (arh *AdminStudentHandler) UpdateStudent(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	students, errs := arh.studentService.UpdateStudent()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(students, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostRole handles POST /v1/admin/roles requests
func (arh *AdminStudentHandler) RegisterStudent(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {


	students, errs := arh.studentService.RegisterStudent()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(students, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteRole handles DELETE /v1/admin/roles/:id requests
func (arh *AdminStudentHandler) DeleteStudent(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	students, errs := arh.studentService.DeleteStudent(1)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(students, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
