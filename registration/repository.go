package registration


import(

	"github.com/solomonkindie/Project/entity"
)


type RegistrationRepository interface{

	RegisterStudent(student entity.Student) error
}