package main

import (
	"database/sql"
	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
)

var tmpl = template.Must(template.ParseGlob("../web/templates/*"))

func main() {

	dbconn, err := sql.Open("postgres",
		"postgres://postgres:aait@localhost/school?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

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

	homeHandler:= handler.NewHomeHandler(tmpl,profileService)
	loginHandler := handler.NewLoginHandler(tmpl, profileService)
	logoutHandler := handler.NewLogoutHandler(tmpl, profileService)


	fs := http.FileServer(http.Dir("../web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux := http.NewServeMux()



	mux.HandleFunc("/", homeHandler.Home)
	mux.HandleFunc("/logout", logoutHandler.Logout)

	// student handler
	mux.HandleFunc("/student/register", registrationHandler.StudentRegistration)
	mux.HandleFunc("/student/login", loginHandler.Login)
	mux.HandleFunc("/student/course", courseHandler.StudentCourse)
	mux.HandleFunc("/student/notification", notificationHandler.StudentGetNotification)
	mux.HandleFunc("/student/profiles", profileHandler.StudentsProfile)
	mux.HandleFunc("/student/profile", profileHandler.StudentProfile)

	mux.HandleFunc("/student/new", attendanceHandler.StudentFillAttendance)
	mux.HandleFunc("/student/show", attendanceHandler.StudentShowAttendance)
	mux.HandleFunc("/student/check", attendanceHandler.StudentCheckAttendance)



	// family handler

	mux.HandleFunc("/family/register", registrationHandler.FamilyRegistration)

	// teacher handler

	mux.HandleFunc("/teacher/register", registrationHandler.TeacherRegistration)
	mux.HandleFunc("/teacher/notification", notificationHandler.TeacherAddNotification)

	// admin handler

	mux.HandleFunc("/admin/register", registrationHandler.AdminRegistration)
	mux.HandleFunc("/admin/student", profileHandler.AdminGetStudent)
	mux.HandleFunc("/admin/student/delete", profileHandler.AdminDeleteStudent)
	mux.HandleFunc("/admin/teacher", profileHandler.AdminGetTeacher)
	mux.HandleFunc("/admin/teacher/delete", profileHandler.AdminDeleteTeacher)
	mux.HandleFunc("/admin/course", courseHandler.AdminGetCourse)
	mux.HandleFunc("/admin/course/new", courseHandler.AdminAddCourse)



	// api hadlers

	mux.HandleFunc("/api/student/course", courseHandler.ApiStudentGetCourse)
	mux.HandleFunc("/api/student/courses", courseHandler.ApiStudentGetCourses)
	mux.HandleFunc("/api/student/notification", notificationHandler.ApiStudentGetNotification)
	mux.HandleFunc("/api/teacher/notification", notificationHandler.TeacherPostNotification)
	mux.HandleFunc("/api/admin/course/add", courseHandler.ApiAdminPostCourse)
	mux.HandleFunc("/api/admin/courses", courseHandler.ApiAdminGetCourses)
	mux.HandleFunc("/api/admin/course/delete", courseHandler.ApiAdminDeleteCourse)


	_ = http.ListenAndServe(":8181", mux)
}
