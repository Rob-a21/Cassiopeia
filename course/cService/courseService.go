package cService



import (
	"github.com/robi_a21/Cassiopeia/course"
	"github.com/robi_a21/Cassiopeia/entity"
)



type CourseServiceImpl struct {
	courseRepository course.CourseRepository
}

func NewCourseServiceImpl(regRepo course.CourseRepository) *CourseServiceImpl {
	
	return &CourseServiceImpl{courseRepository: regRepo}
}

func (ss *CourseServiceImpl) AddCourse(course entity.Course) error {

	err := ss.courseRepository.AddCourse(course)

	if err != nil {
		return err
	}

	return nil
}

func (crs *CourseServiceImpl) GetCourse() ([]entity.Course, error) {

	course, err := crs.courseRepository.GetCourse()

	if err != nil {
		return nil, err
	}

	return course, nil
}

