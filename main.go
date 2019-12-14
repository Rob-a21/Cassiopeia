package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/solomonkindie/Project/delivery/handler"
	"github.com/solomonkindie/Project/registration/repository"
	"github.com/solomonkindie/Project/registration/service"
	"github.com/solomonkindie/Project/profile/pRepository"
	"github.com/solomonkindie/Project/profile/pService"



)

func main() {

	dbconn, err := sql.Open("postgres", "postgres://postgres:aait@127.0.0.1/school?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("delivery/web/templates/*.html"))

	registrationRepository := repository.NewPsqlRegistrationRepositoryImpl(dbconn)
	registrationService := service.NewRegistrationServiceImpl(registrationRepository)
   
	studentRegHandler := handler.NewStudentRegistrationHandler(tmpl, registrationService)

	
	 profileRepository := pRepository.NewPsqlProfileRepositoryImpl(dbconn)
	 profileService := pService.NewProfileServiceImpl(profileRepository)

	 profileHandler := handler.NewProfileHandler(tmpl, profileService)

	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	

	http.HandleFunc("/student", studentRegHandler.StudentRegistrationNew)

	http.HandleFunc("/profile", profileHandler.StudentsProfile)

	http.ListenAndServe(":80", nil)
}
