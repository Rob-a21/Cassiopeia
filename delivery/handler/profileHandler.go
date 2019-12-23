<<<<<<< HEAD
=======



>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
package handler

import(
    "html/template"
	"net/http"
<<<<<<< HEAD
	"github.com/Rob_a21/Cassiopeia/profile"
=======
	"github.com/robi_a21/Cassiopeia/profile"
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae

)


type ProfileHandler struct {
	tmpl        *template.Template
	profileService profile.ProfileService
}

func NewProfileHandler(T *template.Template, PS profile.ProfileService) *ProfileHandler{
	return &ProfileHandler{tmpl: T, profileService: PS}
}

<<<<<<< HEAD
=======

>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
func (prf *ProfileHandler) StudentsProfile(w http.ResponseWriter, r *http.Request) {
	students, err := prf.profileService.Students()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "studentProfile.html", students)
<<<<<<< HEAD
}
=======
}
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
