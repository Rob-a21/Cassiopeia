package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Rob-a21/Cassiopeia/course/cRepository"
	"github.com/Rob-a21/Cassiopeia/course/cService"
	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
	"github.com/Rob-a21/Cassiopeia/notification/nRepository"
	"github.com/Rob-a21/Cassiopeia/notification/nService"
	"github.com/Rob-a21/Cassiopeia/profile/pRepository"
	"github.com/Rob-a21/Cassiopeia/profile/pService"
	"github.com/Rob-a21/Cassiopeia/registration/repository"
	"github.com/Rob-a21/Cassiopeia/registration/service"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))

func main() {

	dbconn, err := sql.Open("postgres",
		"postgres://postgres:strafael@localhost/cassiopeia?sslmode=disable")

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

	profileRepository := pRepository.NewPsqlProfileRepositoryImpl(dbconn)
	profileService := pService.NewProfileServiceImpl(profileRepository)
	profileHandler := handler.NewProfileHandler(tmpl, profileService)

	notificationRepository := nRepository.NewPsqlNotificationRepositoryImpl(dbconn)
	notificationService := nService.NewNotificationServiceImpl(notificationRepository)
	notificationHandler := handler.NewNotificationHandler(tmpl, notificationService)

	courseRepository := cRepository.NewPsqlCourseRepositoryImpl(dbconn)
	courseService := cService.NewCourseServiceImpl(courseRepository)
	courseHandler := handler.NewCourseHandler(tmpl, courseService)

	//homeHandler:= handler.NewHomeHandler(tmpl)
	loginHandler := handler.NewLoginHandler(tmpl, profileService)

	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux := http.NewServeMux()

	//mux.HandleFunc("/", homeHandler.Home)
      // student handler
	mux.HandleFunc("/student/register", registrationHandler.StudentRegistration)
	mux.HandleFunc("/student/login", loginHandler.Login)
	mux.HandleFunc("/student/course", courseHandler.StudentCourse)
	mux.HandleFunc("/student/notification", notificationHandler.StudentGetNotification)
	mux.HandleFunc("/student/profiles", profileHandler.StudentsProfile)
	mux.HandleFunc("/student/profile", profileHandler.StudentProfile)

	// family handler

	mux.HandleFunc("/family/register", registrationHandler.FamilyRegistration)
	//mux.HandleFunc("/family/course", courseHandler.FamilyGetCourse)

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


	_ = http.ListenAndServe(":2121", mux)
}
