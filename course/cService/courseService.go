package cService

import (
	"github.com/Rob-a21/Cassiopeia/course"
	"github.com/Rob-a21/Cassiopeia/entity"
)

type CourseServiceImpl struct {
	courseRepository course.CourseRepository
}

func NewCourseServiceImpl(crRepo course.CourseRepository) *CourseServiceImpl {

	return &CourseServiceImpl{courseRepository: crRepo}
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

func (cs *CourseServiceImpl) Course(id int) (entity.Course, error) {

	course, err := cs.courseRepository.Course(id)

	if err != nil {
		panic(err)
	}

	return course, nil
}

func (crs *CourseServiceImpl) UpdateCourse(course entity.Course) error {

	err := crs.courseRepository.UpdateCourse(course)

	if err != nil {
		return err
	}

	return nil
}

func (crs *CourseServiceImpl) DeleteCourse(id int) error {

	err := crs.courseRepository.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}
