package assessment

import "github.com/Rob-a21/Cassiopeia/entity"

// CategoryService specifies food menu category services
type GradeService interface {

	//assessment of a single assessment(class students)
	Assessments(grade string) ([]entity.Assessment, error)

	//all subject assessments of one student
	singleStudentAssessments(id string) ([]entity.Assessment, error)

	//Single subject assessment of a assessment
	Assessment(studentid int,subjectId string) ([]entity.Assessment, error)

	UpdateGrade(assessment entity.Assessment) error
	DeleteGrade(id int,subjectId string) error
	DeleteGrades(id int) error
	StoreGrade(assessment entity.Assessment) error
}
