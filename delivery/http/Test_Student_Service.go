package http



import (
	"bytes"
	"github.com/Rob-a21/Cassiopeia/course/cRepository"
	"github.com/Rob-a21/Cassiopeia/course/cService"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Rob-a21/Cassiopeia/attendance/aRepository"
	"github.com/Rob-a21/Cassiopeia/attendance/aService"
	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
	"github.com/Rob-a21/Cassiopeia/entity"
)


func TestStudentCourse(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../templates/*"))

	courseRepo := cRepository.NewPsqlCourseRepositoryImpl(nil)
	courseServ := cService.NewCourseServiceImpl(courseRepo)

	adminCourseHandler := handler.NewCourseHandler(tmpl,courseServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/student/course", adminCourseHandler.AdminGetCourse)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/student/course")
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

func TestStudentPostAttendance(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../templates/*"))

	attendanceRepo := aRepository.NewStudentAttendanceRepositoryImpl(nil)
	attendanceServ := aService.NewStudentAttendanceServiceImpl(attendanceRepo)

	studentAttendanceHandler := handler.NewAttendanceHandler(tmpl,attendanceServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/student/attendance/new", studentAttendanceHandler.ApiStudentPostAttendance)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("studentid", string(entity.AttendanceMock.StudentId))


	resp, err := tc.PostForm(sURL+"/student/course/new", form)
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

	if !bytes.Contains(body, []byte("Mock attendance 01")) {
		t.Errorf("want body to contain %q", body)
	}

}
