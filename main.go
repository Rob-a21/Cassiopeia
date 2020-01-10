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

	mux.HandleFunc("/student/register", registrationHandler.StudentRegistration)
	mux.HandleFunc("/student/login", loginHandler.Login)
	mux.HandleFunc("/student/course", courseHandler.StudentGetCourse)
	mux.HandleFunc("/student/notification", notificationHandler.GetNotification)
	mux.HandleFunc("/student/profiles", profileHandler.StudentsProfile)
	mux.HandleFunc("/student/profile", profileHandler.StudentProfile)

	mux.HandleFunc("/family/register", registrationHandler.FamilyRegistration)
	mux.HandleFunc("/family/course", courseHandler.FamilyGetCourse)


	mux.HandleFunc("/teacher/register", registrationHandler.TeacherRegistration)
	mux.HandleFunc("/teacher/notification", notificationHandler.AddNotification)

	mux.HandleFunc("/admin/register", registrationHandler.AdminRegistration)
	mux.HandleFunc("/admin/student", profileHandler.AdminGetStudent)
	mux.HandleFunc("/admin/student/delete", profileHandler.AdminDeleteStudent)
	mux.HandleFunc("/admin/teacher", profileHandler.AdminGetTeacher)
	mux.HandleFunc("/admin/teacher/delete", profileHandler.AdminDeleteTeacher)
	mux.HandleFunc("/admin/course", courseHandler.AdminGetCourse)
	mux.HandleFunc("/admin/course/new", courseHandler.AdminAddCourse)
	mux.HandleFunc("/admin/course/update", courseHandler.AdminUpdateCourse)
	mux.HandleFunc("/admin/course/delete", courseHandler.AdminDeleteCourse)

	mux.HandleFunc("/api/admin/course", courseHandler.AdminPostCourse)

	_ = http.ListenAndServe(":2121", mux)
}
