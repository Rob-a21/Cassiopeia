package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Rob-a21/Cassiopeia/delivery/handler"
	"github.com/Rob-a21/Cassiopeia/profile/pRepository"
	"github.com/Rob-a21/Cassiopeia/profile/pService"
	"github.com/Rob-a21/Cassiopeia/registration/repository"
	"github.com/Rob-a21/Cassiopeia/registration/service"
)

var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))

func main() {

	dbconn, err := sql.Open("postgres", "postgres://postgres:strafael@127.0.0.1/logindb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	registrationRepository := repository.NewPsqlRegistrationRepositoryImpl(dbconn)
	registrationService := service.NewRegistrationServiceImpl(registrationRepository)

	studentRegHandler := handler.NewStudentRegistrationHandler(templ, registrationService)

	profileRepository := pRepository.NewPsqlProfileRepositoryImpl(dbconn)
	profileService := pService.NewProfileServiceImpl(profileRepository)

	profileHandler := handler.NewProfileHandler(templ, profileService)

	fs := http.FileServer(http.Dir("delivery/web/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux := http.NewServeMux()

	mux.HandleFunc("/student", studentRegHandler.StudentRegistrationNew)
	mux.HandleFunc("/", homeHandler)

	http.ListenAndServe(":2121", mux)

	mux.HandleFunc("/profile", profileHandler.StudentsProfile)

	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "mainpage.html", "Welcome")
}

