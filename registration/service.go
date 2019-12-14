package registration

import(

	"github.com/solomonkindie/Project/entity"
)

type RegistrationService interface{

	  RegisterStudent(student entity.Student)error
}