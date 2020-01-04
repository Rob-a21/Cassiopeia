package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/registration"
)

// StudentService implements menu.RoleService interface
type StudentService struct {
	studentRepo registration.StudentRepository
}

func (st *StudentService) RegisterStudent(student entity.Student) (*entity.Student, []error) {
	rl, errs := st.studentRepo.RegisterStudent(student)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

func (st *StudentService) Students() ([]entity.Student, []error) {
	stud, errs := st.studentRepo.Students()
	if len(errs) > 0 {
		return nil, errs
	}
	return stud, errs
}

func (st *StudentService) Student(id int) (*entity.Student, []error) {
	rl, errs := st.studentRepo.Student(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

func (st *StudentService) UpdateStudent(student *entity.Student) (*entity.Student, []error) {
	r1, errs := st.studentRepo.UpdateStudent(student)
	if errs != nil {
		return nil, errs
	}
	return r1, errs
}

func (st *StudentService) DeleteStudent(id int) (*entity.Student, []error) {
	rl, errs := st.studentRepo.DeleteStudent(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

// NewRoleService  returns new RoleService
func NewStudentService(studentRepo registration.StudentRepository) registration.StudentService {
	return &StudentService{studentRepo: studentRepo}
}





