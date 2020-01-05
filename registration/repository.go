package registration

import (
	"github.com/Rob-a21/Cassiopeia/entity"
)

type RegistrationRepository interface {
	RegisterStudent(student entity.Student) error
	RegisterFamily(family entity.Family) error
	RegisterTeacher(teacher entity.Teacher) error
	RegisterAdmin(admin entity.Admin) error
}
