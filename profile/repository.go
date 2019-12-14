package profile

import(
	
	"github.com/solomonkindie/Project/entity"
)
type ProfileRepository interface{

	 Students() ([]entity.Student ,error)

	 Student(id int) (entity.Student, error)

}