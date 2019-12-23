package registration

import(

	"github.com/Rob_a21/Cassiopeia/entity"
)

//RegistrationServ interface
type RegistrationServ interface{

	  RegisterStudent(student entity.Student)error
}
