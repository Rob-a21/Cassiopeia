package handler

import (
	"html/template"
	"net/http"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type LoginHandler struct {
	tmpl         *template.Template
	loginService models.ProfileService
}

func NewLoginHandler(T *template.Template, PS models.ProfileService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginService: PS}
}

func (slh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		studentUser := entity.Student{}
		studentUser.UserName = r.FormValue("username")
		studentUser.Password = r.FormValue("password")

		familyUser := entity.Family{}
		familyUser.Username = r.FormValue("username")
		familyUser.Password = r.FormValue("password")

		teacherUser := entity.Teacher{}
		teacherUser.UserName = r.FormValue("username")
		teacherUser.Password = r.FormValue("password")

		adminUser := entity.Admin{}
		adminUser.UserName = r.FormValue("username")
		adminUser.Password = r.FormValue("password")

		student, err := slh.loginService.Students()
		family, err := slh.loginService.Families()
		teacher, err := slh.loginService.Teachers()
		admin, err := slh.loginService.Admins()

		if err != nil {

			panic(err)
		}

		for index := range student {

			uname := student[index]
			pass := student[index]

			if uname.UserName == studentUser.UserName && pass.Password == studentUser.Password {

				http.Redirect(w, r, "/student", http.StatusSeeOther)

			} else {
			}
		}

		for f := range family {

			uname := family[f]
			pass := family[f]

			if uname.Username == familyUser.Username && pass.Password == familyUser.Password {

				http.Redirect(w, r, "/family", http.StatusSeeOther)
			}
		}

		for t := range teacher {

			uname := teacher[t]
			pass := teacher[t]

			if uname.UserName == teacherUser.UserName && pass.Password == teacherUser.Password {

				http.Redirect(w, r, "/teacher", http.StatusSeeOther)
			}
		}
		for a := range admin {

			uname := admin[a]
			pass := admin[a]

			if uname.UserName == adminUser.UserName && pass.Password == adminUser.Password {

				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			}
		}

	} else {

		slh.tmpl.ExecuteTemplate(w, "login.html", nil)

	}

}
