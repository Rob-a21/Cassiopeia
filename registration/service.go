package registration

import(

	"github.com/robi_a21/Cassiopeia/entity"
)

//RegistrationServ interface
type RegistrationServ interface{

	  RegisterStudent(student entity.Student)error
}
