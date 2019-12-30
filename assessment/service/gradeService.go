package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/assessment"
)

// GradeServiceImpl implements assessment.GradeService interface
type GradeServiceImpl struct {
	GradeRepo assessment.GradeRepository
}

// NewGradeServiceImpl will create new GradeService object
func NewGradeServiceImpl(grRepo assessment.GradeRepository) *GradeServiceImpl {
	return &GradeServiceImpl{GradeRepo: grRepo}
}

// StoreGrade holds new assessment information
func (cs *GradeServiceImpl) StoreGrade(grade entity.Assessment) error {

	err := cs.GradeRepo.StoreGrade(grade)

	if err != nil {
		return err
	}

	return nil
}

// Assessment returns list of assessments
func (cs *GradeServiceImpl) Assessment(studentId int, subjectId string) ([]entity.Assessment, error) {

	grades, err := cs.GradeRepo.Assessment(studentId,subjectId)

	if err != nil {
		return nil, err
	}

	return grades, nil
}

// singleStudentAssessments returns list of categories
func (cs *GradeServiceImpl) singleStudentAssessments(id string) ([]entity.Assessment, error) {

	grades, err := cs.GradeRepo.SingleStudentAssessments(id)

	if err != nil {
		return nil, err
	}

	return grades, nil
}

// Categories returns list of categories
func (cs *GradeServiceImpl) Assessments(grade string) ([]entity.Assessment, error) {

	grades, err := cs.GradeRepo.Assessments(grade)

	if err != nil {
		return nil, err
	}

	return grades, nil
}


// UpdateGrade updates an assessment with new data
func (cs *GradeServiceImpl) UpdateGrade(grade entity.Assessment) error {

	err := cs.GradeRepo.UpdateGrade(grade)

	if err != nil {
		return err
	}

	return nil
}

// DeleteGrade deletes an assessment by the student id and  subject id
func (cs *GradeServiceImpl) DeleteGrade(studentId int, subjectId string) error {

	err := cs.GradeRepo.DeleteGrade(studentId,subjectId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteGrades deletes assessments by the student id
func (cs *GradeServiceImpl) DeleteGrades(studentId int, subjectId string) error {

	err := cs.GradeRepo.DeleteGrade(studentId,subjectId)
	if err != nil {
		return err
	}
	return nil
}

