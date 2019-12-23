package service


import (

	"github.com/Rob_a21/Cassiopeia/entity"
	"github.com/Rob_a21/Cassiopeia/registration"

)



type RegistrationServiceImpl struct {
	registrationRepo registration.RegistrationRepo
}

func NewRegistrationServiceImpl(regRepo registration.RegistrationRepo) *RegistrationServiceImpl {
	
	return &RegistrationServiceImpl{registrationRepo: regRepo}
}

func (ss *RegistrationServiceImpl) RegisterStudent(student entity.Student) error {

	err := ss.registrationRepo.RegisterStudent(student)

	if err != nil {
		return err
	}

	return nil
}