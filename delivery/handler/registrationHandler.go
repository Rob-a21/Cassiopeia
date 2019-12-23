package handler

import(
    "html/template"
	"net/http"
	"strconv"
<<<<<<< HEAD
	"github.com/Rob_a21/Cassiopeia/entity"
	"github.com/Rob_a21/Cassiopeia/registration"
=======
	"github.com/robi_a21/Cassiopeia/entity"
	"github.com/robi_a21/Cassiopeia/registration"
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae

)


type StudentRegistrationHandler struct {
	tmpl        *template.Template
<<<<<<< HEAD
	regService registration.RegistrationServ
}

func NewStudentRegistrationHandler(T *template.Template, RS registration.RegistrationServ) *StudentRegistrationHandler {
=======
	regService registration.RegistrationService
}

func NewStudentRegistrationHandler(T *template.Template, RS registration.RegistrationService) *StudentRegistrationHandler {
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
	return &StudentRegistrationHandler{tmpl: T, regService: RS}
}



func (srh *StudentRegistrationHandler) StudentRegistrationNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		student := entity.Student{}
		student.UserName = r.FormValue("username")
		student.Password = r.FormValue("password")
		student.FirstName = r.FormValue("fname")
		student.LastName = r.FormValue("lname")
		student.ID,_ = strconv.Atoi(r.FormValue("id"))
		student.Email = r.FormValue("email")


		// mf, fh, err := r.FormFile("catimg")
		// if err != nil {
		// 	panic(err)
		// }
		// defer mf.Close()

		// student.Image = fh.Filename

		// writeFile(&mf, fh.Filename)

		srh.regService.RegisterStudent(student)

		// if err != nil {
		// 	panic(err)
		// }

		http.Redirect(w, r, "/student", http.StatusSeeOther)

	} else {

		srh.tmpl.ExecuteTemplate(w, "studentRegNew.html", nil)

	}
}

// func writeFile(mf *multipart.File, fname string) {

// 	wd, err := os.Getwd()

// 	if err != nil {
// 		panic(err)
// 	}

// 	path := filepath.Join(wd, "web", "assets", "img", fname)
// 	image, err := os.Create(path)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer image.Close()
// 	io.Copy(image, *mf)
// }