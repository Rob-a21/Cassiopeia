package registration


import(

	"github.com/Rob_a21/Cassiopeia/entity"
)


type RegistrationRepo interface{

	RegisterStudent(student entity.Student) error
}
