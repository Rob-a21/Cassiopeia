package profile

import (
	"github.com/Rob-a21/Cassiopeia/entity"
)

type ProfileRepository interface {
	Students() ([]entity.Student, error)
	Student(id int) (entity.Student, error)
	DeleteStudent(id int) error
	Families() ([]entity.Family, error)
	Teachers() ([]entity.Teacher, error)
	Teacher(id string) (entity.Teacher, error)
	DeleteTeacher(id string) error
	Admins() ([]entity.Admin, error)
}
