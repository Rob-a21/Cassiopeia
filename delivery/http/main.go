package main


import (
	urepim "github.com/Rob-a21/Cassiopeia/registration/repository"
	usrvim "github.com/Rob-a21/Cassiopeia/registration/service"
	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

dbconn, err := gorm.Open("postgres",
	"postgres://postgres:strafael@127.0.0.1/logindb?sslmode=disable")

if err != nil {
panic(err)
}

defer dbconn.Close()

studentRepo := urepim.NewStudentGormRepoImpl(dbconn)
studentsrv := usrvim.NewStudentService(studentRepo)
adminStudentHandler := handler.NewAdminStudentHandler(studentsrv)

//commentRepo := repository.NewCommentGormRepo(dbconn)
//commentSrv := service.NewCommentService(commentRepo)
//adminCommentHandler := handler.NewAdminCommentHandler(commentSrv)

router := httprouter.New()

router.GET("/v1/main/students", adminStudentHandler.GetStudents)

router.GET("/v1/admin/student/:id", adminStudentHandler.GetSingleStudent)
router.PUT("/v1/admin/students/:id", adminStudentHandler.UpdateStudent)
router.POST("/v1/admin/student", adminStudentHandler.RegisterStudent)
router.DELETE("/v1/admin/students/:id", adminStudentHandler.DeleteStudent)

	_ = http.ListenAndServe(":8181", router)
}

//package main
//
//import (
//	"database/sql"
//	"html/template"
//	"net/http"
//
//	_ "github.com/lib/pq"
//
//	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
//	"github.com/Rob-a21/Cassiopeia/profile/pRepository"
//	"github.com/Rob-a21/Cassiopeia/profile/pService"
//	"github.com/Rob-a21/Cassiopeia/registration/repository"
//	"github.com/Rob-a21/Cassiopeia/registration/service"
//)
//
//var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))
//
//func main() {
//
//	dbconn, err := sql.Open("postgres", "postgres://postgres:strafael@127.0.0.1/logindb?sslmode=disable")
//
//	if err != nil {
//		panic(err)
//	}
//
//	defer dbconn.Close()
//
//	if err := dbconn.Ping(); err != nil {
//		panic(err)
//	}
//
//	registrationRepository := repository.NewPsqlRegistrationRepositoryImpl(dbconn)
//	registrationService := service.NewRegistrationServiceImpl(registrationRepository)
//
//	studentRegHandler := handler.NewStudentRegistrationHandler(templ, registrationService)
//
//	profileRepository := pRepository.NewPsqlProfileRepositoryImpl(dbconn)
//	profileService := pService.NewProfileServiceImpl(profileRepository)
//
//
//	//Handlers
//	profileHandler := handler.NewProfileHandler(templ, profileService)
//
//
//
//	fs := http.FileServer(http.Dir("delivery/web/assets/"))
//	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
//
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/student", studentRegHandler.StudentRegistrationNew)
//	mux.HandleFunc("/", homeHandler)
//	mux.HandleFunc("/Studentprofile", profileHandler.StudentsProfile)
//
//	_ = http.ListenAndServe(":2121", mux)
//
//}
//
//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	_ = templ.ExecuteTemplate(w, "mainpage.html", "Welcome")
//}
//func loginHandler(w http.ResponseWriter, r *http.Request){
//	_ = templ.ExecuteTemplate(w, "login.html", "Welcome")
//}
//func signupHandler(w http.ResponseWriter, r *http.Request){
//	_ = templ.ExecuteTemplate(w, "signup.html", "Welcome")
//}
