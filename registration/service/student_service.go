package service

import (
	"github.com/robi_a21/Cassiopeia/registration"
)

// StudentService implements menu.RoleService interface
type StudentService struct {
	studentRepo registration.RegistrationRepository
}
