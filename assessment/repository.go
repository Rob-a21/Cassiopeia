package assessment

import "github.com/Rob-a21/Cassiopeia/entity"

type AssessmentRepository interface {
	Assessments(grade string) ([]entity.Assessment,error)
	SingleStudentAssessment(id int) ([]entity.Assessment, error)
	Assessment(assessment entity.Assessment) ([]entity.Assessment,error)
	UpdateGrade(assessment entity.Assessment) error
	DeleteGrade(studentID int, subjectID int) error
	DeleteGrades(studentID int) error
	StoreGrade(assessment entity.Assessment) error
	IsQualified(studentID int) (bool, error)
}

