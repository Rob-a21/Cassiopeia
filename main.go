package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/robi_a21/Cassiopeia/course/cRepository"
	"github.com/robi_a21/Cassiopeia/course/cService"
	"github.com/robi_a21/Cassiopeia/delivery/http/handler"
	"github.com/robi_a21/Cassiopeia/notification/nRepository"
	"github.com/robi_a21/Cassiopeia/notification/nService"
	"github.com/robi_a21/Cassiopeia/profile/pRepository"
	"github.com/robi_a21/Cassiopeia/profile/pService"
	"github.com/robi_a21/Cassiopeia/registration/repository"
	"github.com/robi_a21/Cassiopeia/registration/service"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))

func main() {

	dbconn, err := sql.Open("postgres", "postgres://postgres:aait@127.0.0.1/school?sslmode=disable")

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

	mux.HandleFunc("/student/login", loginHandler.Login)
	mux.HandleFunc("/student/course", courseHandler.GetCourse)
	mux.HandleFunc("/student/notification", notificationHandler.GetNotification)
	mux.HandleFunc("/student/register", registrationHandler.StudentRegistration)
	mux.HandleFunc("student/profiles", profileHandler.StudentsProfile)
	mux.HandleFunc("student/profile", profileHandler.StudentProfile)

	mux.HandleFunc("/family/register", registrationHandler.FamilyRegistration)

	mux.HandleFunc("/teacher/register", registrationHandler.TeacherRegistration)
	mux.HandleFunc("/teacher/notification", notificationHandler.AddNotification)

	mux.HandleFunc("/admin/register", registrationHandler.AdminRegistration)
	mux.HandleFunc("/admin/course", courseHandler.CourseAdd)

	http.ListenAndServe(":8181", mux)
}
