package main

import (
	"bytes"
	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
	"github.com/Rob-a21/Cassiopeia/entity"
)

//
//func TestAdminCourse(t *testing.T) {
//
//	tmpl := template.Must(template.ParseGlob("../web/templates/*"))
//
//	courseRepo := repository.NewPsqlCourseRepositoryImpl(nil)
//	courseServ := service.NewCourseServiceImpl(courseRepo)
//
//	adminCourseHandler := handler.NewCourseHandler(tmpl,courseServ)
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/admin/course", adminCourseHandler.ApiAdminGetCourses)
//	ts := httptest.NewTLSServer(mux)
//	defer ts.Close()
//
//	tc := ts.Client()
//	url := ts.URL
//
//	resp, err := tc.Get(url + "/admin/course")
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
//	if !bytes.Contains(body, []byte("Mock course 01")) {
//		t.Errorf("want body to contain %q", body)
//	}
//
//}

func TestAdminAddCourse(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../web/templates/*"))

	courseRepo := repository.NewPsqlCourseRepositoryImpl(nil)
	courseServ := service.NewCourseServiceImpl(courseRepo)

	adminCourseHandler := handler.NewCourseHandler(tmpl,courseServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/course/new", adminCourseHandler.AdminAddCourse)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("coursename", entity.CourseMock.CourseName)
	form.Add("courseid", string(entity.CourseMock.CourseID))
	form.Add("grade", string(entity.CourseMock.Grade))


	resp, err := tc.PostForm(sURL+"/admin/course/new", form)
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

	if !bytes.Contains(body, []byte("Mock Category 01")) {
		t.Errorf("want body to contain %q", body)
	}

}
