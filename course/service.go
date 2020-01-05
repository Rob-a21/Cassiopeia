package course

import (
	"github.com/Rob-a21/Cassiopeia/entity"
)

type CourseService interface {
	AddCourse(course entity.Course) error
	GetCourse() ([]entity.Course, error)
}
