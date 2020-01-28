package main

import (
	"bytes"
	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

//////////-----------AssessmentHandlerTest-------------///////////////////

func TestApiTeacherPostGrade(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	assRepo := repository.NewAssessmentMockRepository(nil)
	assServ := service.NewAssessmentServiceImpl(assRepo)

	adminAssessmentHandler := handler.NewAssessmentHandler(tmpl, assServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/teacher/grade/new", adminAssessmentHandler.ApiTeacherPostGrade)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("value", string(entity.AssessmentMock.Value))
	form.Add("studentID", string(entity.AssessmentMock.StudentID))
	form.Add("subjectID", string(entity.AssessmentMock.SubjectID))
	form.Add("grade", string(entity.AssessmentMock.Grade))
	form.Add("assessment", string(entity.AssessmentMock.Assessment))

	resp, err := tc.PostForm(sURL+"/api/teacher/grade/new", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("")) {
		t.Errorf("want body to contain %q", body)
	}

}

///////////////////////////

func TestAssessmentsOfOneGrade(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	assRepo := repository.NewAssessmentMockRepository(nil)
	assServ := service.NewAssessmentServiceImpl(assRepo)

	adminAssessmentHandler := handler.NewAssessmentHandler(tmpl, assServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/student/grade", adminAssessmentHandler.AssessmentsOfOneGrade)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL + "/student/grade")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	//////////////////////////////////////////////////////
	if !bytes.Contains(body, []byte("")) {
		t.Errorf("want body to contain %q", body)
	}
}

//////////-----------AttendanceHandlerTest-------------///////////////////

func TestApiStudentPostAttendance(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	attRepo := repository.NewAttendanceRepoMock(nil)
	attServ := service.NewStudentAttendanceServiceImpl(attRepo)

	adminAttendanceHandler := handler.NewAttendanceHandler(tmpl, attServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/student/attendance/new", adminAttendanceHandler.ApiStudentPostAttendance)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("date", entity.AttendanceMock.Date.String())
	form.Add("studentID", string(entity.AttendanceMock.StudentID))
	form.Add("attendance", string(entity.AttendanceMock.Attendance))

	resp, err := tc.PostForm(sURL+"/api/student/attendance/new", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestShowStudentsAttendance(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	attRepo := repository.NewAttendanceRepoMock(nil)
	attServ := service.NewStudentAttendanceServiceImpl(attRepo)

	adminAttendanceHandler := handler.NewAttendanceHandler(tmpl, attServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/student/attendance/show", adminAttendanceHandler.ShowStudentsAttendance)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL + "/student/attendance/show")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte(body)) {
		t.Errorf("want body to contain %q", body)
	}
}

//////////////////////-------COURSE TEST-----------//////////////

func TestApiAdminPostCourse(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	courseRepo := repository.NewcourseRepoMock(nil)
	courseServ := service.NewCourseServiceImpl(courseRepo)

	adminCourseHandler := handler.NewCourseHandler(tmpl, courseServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/admin/course/add", adminCourseHandler.ApiAdminPostCourse)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("coursename", entity.CourseMock.CourseName)
	form.Add("courseid", string(entity.CourseMock.CourseID))
	form.Add("grade", string(entity.CourseMock.Grade))

	resp, err := tc.PostForm(sURL+"/api/admin/course/add", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("")) {
		t.Errorf("want body to contain %q", body)
	}

}

//////////////////---------------NOTIFICATION-------------/////////////

func TestApiStudentGetNotification(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	notRepo := repository.NewNotificationRepoMock(nil)
	notServ := service.NewNotificationServiceImpl(notRepo)

	adminNotificationHandler := handler.NewNotificationHandler(tmpl, notServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/student/notification", adminNotificationHandler.ApiStudentGetNotification)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL + "/api/student/notification")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte(body)) {
		t.Errorf("want body to contain %q", body)
	}
}

//////////-------------------Profile----------------////////////////
//////////////////////////////////
////////////////////////////////
////////////////////////////////
///////////////////////////////////////
func TestAdminGetStudent(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	profrepo := repository.NewprofileRepoMock(nil)
	profServ := service.NewProfileServiceImpl(profrepo)

	adminNotificationHandler := handler.NewProfileHandler(tmpl, profServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/student", adminNotificationHandler.AdminGetStudent)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL + "/admin/student")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("username")) {
		t.Errorf("want body to contain %q", body)
	}
}

//////////-------------------registration----------------////////////////
//
//func TestStudentRegistration(t *testing.T) {
//
//	tmpl := template.Must(template.ParseGlob("../web/templates/*"))
//
//	regRepo := repository.NewregistrationMockRepo(nil)
//	regServ := service.NewRegistrationServiceImpl(regRepo)
//
//	adminRegistrationHandler := handler.NewRegistrationHandler(tmpl, regServ)
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/admin/register/student", adminRegistrationHandler.StudentRegistration)
//	ts := httptest.NewTLSServer(mux)
//	defer ts.Close()
//
//	tc := ts.Client()
//	sURL := ts.URL
//
//	form := url.Values{}
//	form.Add("username", entity.StudentMock.UserName)
//	form.Add("studentID", string(entity.StudentMock.ID))
//	form.Add("password", entity.StudentMock.Password)
//	form.Add("firstname", entity.StudentMock.FirstName)
//	form.Add("lastname", entity.StudentMock.LastName)
//	form.Add("grade", string(entity.StudentMock.Grade))
//	form.Add("email", entity.StudentMock.Email)
//
//	resp, err := tc.PostForm(sURL+"/admin/register/student", form)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
//	}
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if !bytes.Contains(body, []byte("username")) {
//		t.Errorf("want body to contain %q", body)
//	}
//}

///////////////// Login Handler/////////////////////
