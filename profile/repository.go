package profile

import(
	
	"github.com/robi_a21/Cassiopeia/entity"
)
type ProfileRepository interface{

	 Students() ([]entity.Student ,error)
	 Student(id int) (entity.Student, error)
	 Families() ([]entity.Family ,error)
	 Teachers() ([]entity.Teacher ,error)
	 Admins() ([]entity.Admin ,error)


}