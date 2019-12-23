package profile

import(
	
	"github.com/Rob_a21/Cassiopeia/entity"
)
type ProfileRepository interface{

	 Students() ([]entity.Student ,error)

	 Student(id int) (entity.Student, error)

}
