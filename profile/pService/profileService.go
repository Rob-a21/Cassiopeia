package pService

import (
	"github.com/solomonkindie/Project/entity"
	"github.com/solomonkindie/Project/profile"
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


