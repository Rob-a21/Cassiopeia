package main

import (
	"database/sql"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"
	"github.com/Rob-a21/Cassiopeia/token"
	"time"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
)


func main() {

	var tmpl = template.Must(template.ParseGlob("../web/templates/*"))

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


	homeHandler:= handler.NewHomeHandler(tmpl,profileService)
	loginHandler := handler.NewLoginHandler(tmpl, profileService)
	logoutHandler := handler.NewLogoutHandler(tmpl, profileService)

	//sess := configSess()
	//uh := handler.RegistrationHandler(tmpl, registrationService, sessionSrv, sess, csrfSignKey)



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

	mux.HandleFunc("/student/new", attendanceHandler.FillStudentAttendance)
	mux.HandleFunc("/student/show", attendanceHandler.ShowStudentsAttendance)
	mux.HandleFunc("/student/check", attendanceHandler.CheckStudentAttendance)

	mux.HandleFunc("/student/grade", assessmentHandler.AssessmentsOfOneGrade)



	// family handler

	mux.HandleFunc("/family/register", registrationHandler.FamilyRegistration)

	// teacher handler

	mux.HandleFunc("/teacher/register", registrationHandler.TeacherRegistration)
	mux.HandleFunc("/teacher/notification", notificationHandler.TeacherAddNotification)
	mux.HandleFunc("/teacher/grade/new", assessmentHandler.StoreGrade)
	mux.HandleFunc("/teacher/grade/update", assessmentHandler.UpdateGrade)
	mux.HandleFunc("/teacher/grade/delete", assessmentHandler.DeleteGrade)
	mux.HandleFunc("/teacher/grade/deletes", assessmentHandler.DeleteGrades)




	// admin handler

	mux.HandleFunc("/admin/register", registrationHandler.AdminRegistration)
	mux.HandleFunc("/admin/student", profileHandler.AdminGetStudent)
	mux.HandleFunc("/admin/student/delete", profileHandler.AdminDeleteStudent)
	mux.HandleFunc("/admin/teacher", profileHandler.AdminGetTeacher)
	mux.HandleFunc("/admin/teacher/delete", profileHandler.AdminDeleteTeacher)
	mux.HandleFunc("/admin/course", courseHandler.AdminGetCourse)
	mux.HandleFunc("/admin/course/new", courseHandler.AdminAddCourse)



	// api hadlers

	mux.HandleFunc("/api/student/attendance/new", attendanceHandler.ApiStudentPostAttendance)
	mux.HandleFunc("/api/student/attendance/check", attendanceHandler.ApiStudentCheckAttendance)
	mux.HandleFunc("/api/student/attendance/show", attendanceHandler.ApiStudentShowAttendance)
	mux.HandleFunc("/api/teacher/grade/new", assessmentHandler.ApiTeacherPostGrade)
	mux.HandleFunc("/api/student/course", courseHandler.ApiStudentGetCourse)
	mux.HandleFunc("/api/student/courses", courseHandler.ApiStudentGetCourses)
	mux.HandleFunc("/api/student/notification", notificationHandler.ApiStudentGetNotification)
	mux.HandleFunc("/api/teacher/notification", notificationHandler.TeacherPostNotification)
	mux.HandleFunc("/api/admin/course/add", courseHandler.ApiAdminPostCourse)
	mux.HandleFunc("/api/admin/courses", courseHandler.ApiAdminGetCourses)
	mux.HandleFunc("/api/admin/course/delete", courseHandler.ApiAdminDeleteCourse)


	_ = http.ListenAndServe(":8181", mux)
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