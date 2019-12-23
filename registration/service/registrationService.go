package service


import (

<<<<<<< HEAD
	"github.com/Rob_a21/Cassiopeia/entity"
	"github.com/Rob_a21/Cassiopeia/registration"
=======
	"github.com/robi_a21/Cassiopeia/entity"
	"github.com/robi_a21/Cassiopeia/registration"
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae

)



type RegistrationServiceImpl struct {
<<<<<<< HEAD
	registrationRepo registration.RegistrationRepo
}

func NewRegistrationServiceImpl(regRepo registration.RegistrationRepo) *RegistrationServiceImpl {
=======
	registrationRepo registration.RegistrationRepository
}

func NewRegistrationServiceImpl(regRepo registration.RegistrationRepository) *RegistrationServiceImpl {
>>>>>>> ef863e83e75485d3bbb1e4923fc31937ecb7d8ae
	
	return &RegistrationServiceImpl{registrationRepo: regRepo}
}

func (ss *RegistrationServiceImpl) RegisterStudent(student entity.Student) error {

	err := ss.registrationRepo.RegisterStudent(student)

	if err != nil {
		return err
	}

	return nil
}