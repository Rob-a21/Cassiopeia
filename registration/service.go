package registration

import "github.com/Rob-a21/Cassiopeia/entity"

// CategoryService specifies food menu category services
type StudentService interface {

	RegisterStudent(student entity.Student) (*entity.Student, []error)
	Students() ([]entity.Student, []error)
	Student(id int) (*entity.Student, []error)
	UpdateStudent(student *entity.Student) (*entity.Student, []error)
	DeleteStudent(id int) (*entity.Student, []error)
}




