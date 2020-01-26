package handler

import (
	"encoding/json"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"
)


type AttendanceHandler struct {
	tmpl              *template.Template
	attendanceService models.StudentAttendanceService
}

func NewAttendanceHandler(T *template.Template, NS models.StudentAttendanceService) *AttendanceHandler {
	return &AttendanceHandler{tmpl: T, attendanceService: NS}
}

func (at *AttendanceHandler) FillStudentAttendance(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		attendance := entity.Attendance{}
		attendance.Date = time.Now()
		attendance.StudentID,_ = strconv.Atoi(r.FormValue("id"))

		_ = at.attendanceService.FillAttendance(attendance)


		http.Redirect(w,r,"/student",http.StatusSeeOther)

	}

	_ = at.tmpl.ExecuteTemplate(w, "student.attendance.new.layout", nil)

}


func (at *AttendanceHandler) CheckStudentAttendance(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("studentid")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		attendance, err := at.attendanceService.CheckAttendance(id)

		if err != nil {
			panic(err)
		}

		at.tmpl.ExecuteTemplate(w, "admin.course.update.layout", attendance)

	}

	}

func (at *AttendanceHandler) ShowStudentsAttendance(w http.ResponseWriter, r *http.Request) {

	attendances, err := at.attendanceService.ShowAttendance()
	if err != nil {
		panic(err)
	}

	at.tmpl.ExecuteTemplate(w, "student.attendance.layout", attendances)


	}





func (at *AttendanceHandler)ApiStudentPostAttendance(w http.ResponseWriter,r *http.Request){

	len := r.ContentLength

	body:= make([]byte,len)

	r.Body.Read(body)

	attendance:= entity.Attendance{}

	json.Unmarshal(body,&attendance)


	at.attendanceService.FillAttendance(attendance)

	w.WriteHeader(200)

	return
}







func (ntf *AttendanceHandler)ApiStudentCheckAttendance(w http.ResponseWriter,r *http.Request) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil{

		return
	}


	attendance := entity.Attendance{}

	_, _ = ntf.attendanceService.CheckAttendance(id)

	output,err := json.MarshalIndent(&attendance,"","\t\t")

	if err != nil{

		return
	}

	w.Header().Set("Content-Type","application/json")

	w.Write(output)

	return
}
//
func (att *AttendanceHandler)ApiStudentShowAttendance(w http.ResponseWriter,r *http.Request){

	len := r.ContentLength

	body:= make([]byte,len)

	r.Body.Read(body)

	attendance:= entity.Attendance{}

	json.Unmarshal(body,&attendance)


	att.attendanceService.ShowAttendance()

	w.WriteHeader(200)

	return
}
