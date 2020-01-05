package handler

import (
	"html/template"
	"net/http"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/profile"
)

type LoginHandler struct {
	tmpl         *template.Template
	loginService profile.ProfileService
}

func NewLoginHandler(T *template.Template, PS profile.ProfileService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginService: PS}
}

func (srh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		studentUser := entity.Student{}
		studentUser.UserName = r.FormValue("username")
		studentUser.Password = r.FormValue("password")

		// familyUser := entity.Family{}
		// familyUser.Username = r.FormValue("username")
		// familyUser.Password = r.FormValue("password")

		// teacherUser := entity.Teacher{}
		// teacherUser.UserName = r.FormValue("username")
		// teacherUser.Password = r.FormValue("password")

		// adminUser := entity.Admin{}
		// adminUser.UserName = r.FormValue("username")
		// adminUser.Password = r.FormValue("password")

		student, err := srh.loginService.Students()
		// family, err := srh.loginService.Families()
		// teacher, err := srh.loginService.Teachers()
		// admin, err := srh.loginService.Admins()

		if err != nil {

			panic(err)
		}

		for s := range student {

			uname := student[s]
			pass := student[s]

			if uname.UserName == studentUser.UserName && pass.Password == studentUser.Password {

				http.Redirect(w, r, "/student/profiles", http.StatusSeeOther)

			} else {
			}
		}

		// for f := range family {

		// 	uname := family[f]
		// 	pass := family[f]

		// 	if uname.Username == familyUser.Username && pass.Password == familyUser.Password {

		// 		http.Redirect(w, r, "", http.StatusSeeOther)
		// 	}
		// }

		// for t := range teacher {

		// 	uname := teacher[t]
		// 	pass := teacher[t]

		// 	if uname.UserName == teacherUser.UserName && pass.Password == teacherUser.Password {

		// 		http.Redirect(w, r, "teacher/profiles", http.StatusSeeOther)
		// 	}

		// 	for a := range admin {

		// 		uname := admin[a]
		// 		pass := admin[a]

		// 		if uname.UserName == adminUser.UserName && pass.Password == adminUser.Password {

		// 			http.Redirect(w, r, "", http.StatusSeeOther)
		// 		}
		// 	}

		// }

	} else {

		srh.tmpl.ExecuteTemplate(w, "login.html", nil)

	}
}
