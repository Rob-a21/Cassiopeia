


package handler

import(
    "html/template"
	"net/http"
	//"strconv"
	//"github.com/solomonkindie/Project/entity"
	"github.com/solomonkindie/Project/profile"

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
	prf.tmpl.ExecuteTemplate(w, "studentProfile.html", students)
}