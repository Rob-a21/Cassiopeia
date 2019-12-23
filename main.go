package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/robi_a21/Cassiopeia/delivery/handler"
	"github.com/robi_a21/Cassiopeia/profile/pRepository"
	"github.com/robi_a21/Cassiopeia/profile/pService"
	"github.com/robi_a21/Cassiopeia/registration/repository"
	"github.com/robi_a21/Cassiopeia/registration/service"
)

<<<<<<< HEAD
func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "main.layout", "welcome")
}

=======
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
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

	studentRegHandler := handler.NewStudentRegistrationHandler(tmpl, registrationService)

	profileRepository := pRepository.NewPsqlProfileRepositoryImpl(dbconn)
	profileService := pService.NewProfileServiceImpl(profileRepository)

	profileHandler := handler.NewProfileHandler(tmpl, profileService)

	fs := http.FileServer(http.Dir("delivery/web/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux := http.NewServeMux()

	mux.HandleFunc("/student", studentRegHandler.StudentRegistrationNew)
	mux.HandleFunc("/", homeHandler)
<<<<<<< HEAD
	http.ListenAndServe(":2121", mux)

=======
	mux.HandleFunc("/profile", profileHandler.StudentsProfile)

	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "mainpage.html", "Welcome")
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
}

