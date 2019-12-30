package service


import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/registration"

)

type RegistrationServiceImpl struct {
	registrationRepo registration.StudentRepository
}

func NewRegistrationServiceImpl(regRepo registration.StudentRepository) *RegistrationServiceImpl {
	return &RegistrationServiceImpl{registrationRepo: regRepo}
}

func (ss *RegistrationServiceImpl) RegisterStudent(student *entity.Student) (*entity.Student, []error) {

	s1, err := ss.registrationRepo.RegisterStudent(student)

	if err != nil {
		return nil, err
	}

	return s1, err
}

func (ss *RegistrationServiceImpl) Students() ([]entity.Student, []error) {

	students ,err := ss.registrationRepo.Students()

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (ss *RegistrationServiceImpl) Student(id int) (*entity.Student, []error) {

	student ,err := ss.registrationRepo.Student(id)

	if err != nil {
		return student, err
	}

	return student, nil
}


func (ss *RegistrationServiceImpl) UpdateStudent(student *entity.Student) (*entity.Student, []error) {

	s1, err := ss.registrationRepo.UpdateStudent(student)

	if err != nil {
		return nil,err
	}

	return s1, nil
}
func (ss *RegistrationServiceImpl) DeleteStudent(id int) (*entity.Student, []error) {

	s1, err := ss.registrationRepo.DeleteStudent(id)

	if err != nil {
		return nil,err
	}

	return s1, nil
}
