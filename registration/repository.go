package registration


import(

	"github.com/robi_a21/Cassiopeia/entity"
)


type RegistrationRepository interface{

	RegisterStudent(student entity.Student) error
}