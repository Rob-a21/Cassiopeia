package service


import (

	"github.com/solomonkindie/Project/entity"
	"github.com/solomonkindie/Project/registration"

)



type RegistrationServiceImpl struct {
	registrationRepo registration.RegistrationRepository
}

func NewRegistrationServiceImpl(regRepo registration.RegistrationRepository) *RegistrationServiceImpl {
	
	return &RegistrationServiceImpl{registrationRepo: regRepo}
}

func (ss *RegistrationServiceImpl) RegisterStudent(student entity.Student) error {

	err := ss.registrationRepo.RegisterStudent(student)

	if err != nil {
		return err
	}

	return nil
}