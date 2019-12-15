package service


import (

	"github.com/robi_a21/Cassiopeia/entity"
	"github.com/robi_a21/Cassiopeia/registration"

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