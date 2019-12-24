package assessment

import "github.com/Rob-a21/Cassiopeia/entity"

// Grade specifies assessment related database operations
type GradeRepository interface {

	//assessment of a single assessment(one class students)
	Assessments(grade string) ([]entity.Assessment, error)

	//all subject assessments of one student
	SingleStudentAssessments(id string) ([]entity.Assessment, error)

	//Single subject assessment of a class
	Assessment(studentid int,subjectId string) ([]entity.Assessment, error)

	UpdateGrade(assessment entity.Assessment) error
	DeleteGrade(id int,subjectId string) error
	DeleteGrades(id int) error
	StoreGrade(assessment entity.Assessment) error
}

