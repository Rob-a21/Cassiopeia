package repository

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/registration"
	"github.com/jinzhu/gorm"
)

type StudentGormRepo struct {
	conn *gorm.DB
}

// NewStudentRepositoryImpl returns new StudentRepositoryImpl object
func NewStudentGormRepoImpl(db *gorm.DB) registration.StudentRepository {
	return &StudentGormRepo{conn:db}
}

func (st *StudentGormRepo) RegisterStudent(student *entity.Student) (*entity.Student, []error) {
	s := student
	errs := StudentGormRepo{}.conn.Create(s).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return s, nil
}


func (st *StudentGormRepo) Students() ([]entity.Student, []error) {
	student := []entity.Student{}
	errs := st.conn.Find(&student).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return student, errs
}

// Student retrieves a student by its id from the database
func (st *StudentGormRepo) Student(id int) (*entity.Student, []error) {
	student := entity.Student{}
	errs := StudentGormRepo.conn.First(&student, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &student, errs
}

func (st *StudentGormRepo) UpdateStudent(student *entity.Student) (*entity.Student, []error) {
	s := student
	errs := StudentGormRepo{}.conn.Save(s).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return s, nil
}

func (st *StudentGormRepo) DeleteStudent(id int) (*entity.Student, []error) {
	s, errs := StudentGormRepo.Student(*st,id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = StudentGormRepo.conn.Delete(s, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return s, nil
}


//// Role retrieves a role by its id from the database

//
//// UpdateRole updates a given user role in the database
//func (roleRepo *RoleGormRepo) UpdateRole(role *entity.Role) (*entity.Role, []error) {
//	r := role
//	errs := roleRepo.conn.Save(r).GetErrors()
//	if len(errs) > 0 {
//		return nil, errs
//	}
//	return r, errs
//}
//
//// DeleteRole deletes a given user role from the database
//func (roleRepo *RoleGormRepo) DeleteRole(id uint) (*entity.Role, []error) {
//	r, errs := roleRepo.Role(id)
//	if len(errs) > 0 {
//		return nil, errs
//	}
//	errs = roleRepo.conn.Delete(r, id).GetErrors()
//	if len(errs) > 0 {
//		return nil, errs
//	}
//	return r, errs
//}
//
//// StoreRole stores a given user role in the database
//func (roleRepo *RoleGormRepo) StoreRole(role *entity.Role) (*entity.Role, []error) {
//	r := role
//	errs := roleRepo.conn.Create(r).GetErrors()
//	if len(errs) > 0 {
//		return nil, errs
//	}
//	return r, errs
//}
//
