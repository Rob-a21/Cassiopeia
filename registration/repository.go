package registration


import(

	"github.com/robi_a21/Cassiopeia/entity"
)


type RegistrationRepo interface{

	RegisterStudent(student entity.Student) error
}
