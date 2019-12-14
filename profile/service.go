package profile

import(

	"github.com/solomonkindie/Project/entity"
)
type ProfileService interface{

	 Students() ([]entity.Student ,error)

	 Student(id int) (entity.Student, error)
}