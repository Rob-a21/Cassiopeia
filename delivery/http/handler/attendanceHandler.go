package handler

import (
	"encoding/json"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/user"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"
)


type AttendanceHandler struct {
	tmpl                *template.Template
	attendanceService user.StudentAttendanceService
}

func NewAttendanceHandler(T *template.Template, NS user.StudentAttendanceService) *AttendanceHandler {
	return &AttendanceHandler{tmpl: T, attendanceService: NS}
}

func (at *AttendanceHandler) StudentFillAttendance(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		attendance2 := entity.Attendance{}
		attendance2.Date = time.Now()
		attendance2.StudentId,_ = strconv.Atoi(r.FormValue("studentid"))

		_ = at.attendanceService.FillAttendance(attendance2)


	}

	_ = at.tmpl.ExecuteTemplate(w, "admin.course.new.layout", nil)

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


	attendance2 := entity.Attendance{}

	_, _ = ntf.attendanceService.CheckAttendance(id)

	output,err := json.MarshalIndent(&attendance2,"","\t\t")

	if err != nil{

		return
	}

	w.Header().Set("Content-Type","application/json")

	w.Write(output)

	return
}
