package main

import (
	"database/sql"
	"time"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"
	"github.com/Rob-a21/Cassiopeia/token"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
)

var tmpl = template.Must(template.ParseGlob("c:/Users/solki/go/src/github.com/Rob-a21/Cassiopeia/delivery/web/templates/*"))

func main() {

	//csrfSignKey := []byte(token.GenerateRandomID(32))

	dbconn, err := sql.Open("postgres",
		"postgres://postgres:aait@localhost/school?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	//sessionRepo := repository.NewSessionRepo(dbconn)
	//sessionSrv := service.New(sessionRepo)

	registrationRepository := repository.NewPsqlRegistrationRepositoryImpl(dbconn)
	registrationService := service.NewRegistrationServiceImpl(registrationRepository)
	registrationHandler := handler.NewRegistrationHandler(tmpl, registrationService)

	profileRepository := repository.NewPsqlProfileRepositoryImpl(dbconn)
	profileService := service.NewProfileServiceImpl(profileRepository)
	profileHandler := handler.NewProfileHandler(tmpl, profileService)

	notificationRepository := repository.NewPsqlNotificationRepositoryImpl(dbconn)
	notificationService := service.NewNotificationServiceImpl(notificationRepository)
	notificationHandler := handler.NewNotificationHandler(tmpl, notificationService)

	courseRepository := repository.NewPsqlCourseRepositoryImpl(dbconn)
	courseService := service.NewCourseServiceImpl(courseRepository)
	courseHandler := handler.NewCourseHandler(tmpl, courseService)

	attendanceRepository := repository.NewStudentAttendanceRepositoryImpl(dbconn)
	attendanceService := service.NewStudentAttendanceServiceImpl(attendanceRepository)
	attendanceHandler := handler.NewAttendanceHandler(tmpl, attendanceService)

	assessmentRepository := repository.NewAssessmentRepositoryImpl(dbconn)
	assessmentService := service.NewAssessmentServiceImpl(assessmentRepository)
	assessmentHandler := handler.NewAssessmentHandler(tmpl, assessmentService)

	homeHandler := handler.NewHomeHandler(tmpl, profileService)
	loginHandler := handler.NewLoginHandler(tmpl, profileService)
	logoutHandler := handler.NewLogoutHandler(tmpl, profileService)

	//sess := configSess()
	//uh := handler.RegistrationHandler(tmpl, registrationService, sessionSrv, sess, csrfSignKey)

	fs := http.FileServer(http.Dir("c:/Users/solki/go/src/github.com/Rob-a21/Cassiopeia/delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", homeHandler.Home)
	http.HandleFunc("/admin", homeHandler.Admin)
	http.HandleFunc("/student", homeHandler.Student)
	http.HandleFunc("/teacher", homeHandler.Teacher)
	http.HandleFunc("/family", homeHandler.Family)

	http.HandleFunc("/admin/login", loginHandler.AdminLogin)
	http.HandleFunc("/student/login", loginHandler.StudentLogin)
	http.HandleFunc("/teacher/login", loginHandler.TeacherLogin)
	http.HandleFunc("/family/login", loginHandler.FamilyLogin)

	http.HandleFunc("/logout", logoutHandler.Logout)

	// student handler
	http.HandleFunc("/student/course", courseHandler.StudentCourse)
	http.HandleFunc("/student/notification", notificationHandler.StudentGetNotification)
	http.HandleFunc("/student/profiles", profileHandler.StudentsProfile)
	http.HandleFunc("/student/attendance/new", attendanceHandler.FillStudentAttendance)
	http.HandleFunc("/student/attendance/show", attendanceHandler.ShowStudentsAttendance)
	http.HandleFunc("/student/attendance/check", attendanceHandler.CheckStudentAttendance)
	http.HandleFunc("/student/grade", assessmentHandler.AssessmentsOfOneGrade)

	// family handler


	// teacher handler

	http.HandleFunc("/teacher/profile", profileHandler.TeacherProfile)
	http.HandleFunc("/teacher/notification", notificationHandler.TeacherAddNotification)
	http.HandleFunc("/teacher/assessment/new", assessmentHandler.StoreGrade)
	http.HandleFunc("/teacher/assessment/update", assessmentHandler.UpdateGrade)
	http.HandleFunc("/teacher/assessment/delete", assessmentHandler.DeleteGrade)
	http.HandleFunc("/teacher/assessment/deletes", assessmentHandler.DeleteGrades)

	// admin handler

	http.HandleFunc("/admin/register/admin", registrationHandler.AdminRegistration)
	http.HandleFunc("/admin/register/student", registrationHandler.StudentRegistration)
	http.HandleFunc("/admin/register/teacher", registrationHandler.TeacherRegistration)
	http.HandleFunc("/admin/register/family", registrationHandler.FamilyRegistration)
	http.HandleFunc("/admin/profile", profileHandler.AdminProfile)
	http.HandleFunc("/admin/student", profileHandler.AdminGetStudent)
	http.HandleFunc("/admin/student/delete", profileHandler.AdminDeleteStudent)
	http.HandleFunc("/admin/teacher", profileHandler.AdminGetTeacher)
	http.HandleFunc("/admin/teacher/delete", profileHandler.AdminDeleteTeacher)
	http.HandleFunc("/admin/course", courseHandler.AdminGetCourse)
	http.HandleFunc("/admin/course/new", courseHandler.AdminAddCourse)

	// api hadlers

	http.HandleFunc("/api/student/attendance/new", attendanceHandler.ApiStudentPostAttendance)
	http.HandleFunc("/api/student/attendance/check", attendanceHandler.ApiStudentCheckAttendance)
	http.HandleFunc("/api/student/attendance/show", attendanceHandler.ApiStudentShowAttendance)
	http.HandleFunc("/api/teacher/grade/new", assessmentHandler.ApiTeacherPostGrade)
	http.HandleFunc("/api/student/course", courseHandler.ApiStudentGetCourse)
	http.HandleFunc("/api/student/courses", courseHandler.ApiStudentGetCourses)
	http.HandleFunc("/api/student/notification", notificationHandler.ApiStudentGetNotification)
	http.HandleFunc("/api/teacher/notification", notificationHandler.TeacherPostNotification)
	http.HandleFunc("/api/admin/course/add", courseHandler.ApiAdminPostCourse)
	http.HandleFunc("/api/admin/courses", courseHandler.ApiAdminGetCourses)
	http.HandleFunc("/api/admin/course/delete", courseHandler.ApiAdminDeleteCourse)

	http.ListenAndServe(":8181", nil)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := token.GenerateRandomID(32)
	signingString, err := token.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
