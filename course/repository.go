package course

import (
	"github.com/Rob-a21/Cassiopeia/entity"
)

type CourseRepository interface {
	AddCourse(course entity.Course) error
	GetCourse() ([]entity.Course, error)
	Course(id int) (entity.Course, error)
	UpdateCourse(course entity.Course) error
	DeleteCourse(id int) error
}
