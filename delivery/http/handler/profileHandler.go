


package handler

import(
    "html/template"
	"net/http"
	"strconv"
	"github.com/robi_a21/Cassiopeia/entity"
	"github.com/robi_a21/Cassiopeia/profile"


)


type ProfileHandler struct {
	tmpl        *template.Template
	profileService profile.ProfileService
}

func NewProfileHandler(T *template.Template, PS profile.ProfileService) *ProfileHandler{
	return &ProfileHandler{tmpl: T, profileService: PS}
}


func (prf *ProfileHandler) StudentsProfile(w http.ResponseWriter, r *http.Request) {
	
	students, err := prf.profileService.Students()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "student.index.layout", students)

	
}





func (prf *ProfileHandler) StudentProfile(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == http.MethodGet {

		//idRaw := r.URL.Query().Get("id")
		//id, err := strconv.Atoi(idRaw)

		// if err != nil {
		// 	panic(err)
		// }

		st:=entity.Student{}

		st.ID,_ = strconv.Atoi(r.FormValue("id"))

		student, err := prf.profileService.Student(st.ID)

		if err != nil {
			panic(err)
		}

	  //http.Redirect(w,r,"/profile",http.StatusSeeOther)
	  
	  prf.tmpl.ExecuteTemplate(w, "studentProfile.html", student)

	}

	prf.tmpl.ExecuteTemplate(w, "studentByID.html", nil)

}