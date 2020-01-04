package pService

import (
	"github.com/robi_a21/Cassiopeia/entity"
	"github.com/robi_a21/Cassiopeia/profile"
)

type ProfileServiceImpl struct {
	profileRepository profile.ProfileRepository
}

func NewProfileServiceImpl(profrepo profile.ProfileRepository) *ProfileServiceImpl {
	return &ProfileServiceImpl{profileRepository: profrepo}
}

func (ss *ProfileServiceImpl) Students() ([]entity.Student, error) {

	students, err := ss.profileRepository.Students()

	if err != nil {
		return nil, err
	}

	return students, nil
}


func (ss *ProfileServiceImpl) Student(id int) (entity.Student, error) {

	student, err := ss.profileRepository.Student(id)

	if err != nil {
		panic(err)
	}

	return student, nil
}

func (fs *ProfileServiceImpl) Families() ([]entity.Family,error) {

	family, err := fs.profileRepository.Families()

	if err != nil{

		return  nil,err
	}

	return family,nil

}

func (ts *ProfileServiceImpl) Teachers() ([]entity.Teacher,error) {

	teacher, err := ts.profileRepository.Teachers()

	if err != nil{

		return  nil,err
	}

	return teacher,nil

}

func (as *ProfileServiceImpl) Admins() ([]entity.Admin,error){

	admin,err := as.profileRepository.Admins()

	if err != nil{

		return  nil,err
	}

	return  admin,nil
}


