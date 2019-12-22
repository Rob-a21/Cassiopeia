package repository

import "database/sql"
import "github.com/robi_a21/Cassiopeia/entity"
import "errors"

//AssessmentRepositoryImpl struct
type AssessmentRepositoryImpl struct {
	conn *sql.DB
}

//NewAssessmentImpl function
func NewAssessmentImpl(Conn *sql.DB) *AssessmentRepositoryImpl{
	return &AssessmentRepositoryImpl{conn:Conn}
}

//Assessments returned
func (as *AssessmentRepositoryImpl) Assessments( grade string) ([]entity.Assessment, error){
	rows, err := as.conn.Query("SELECT * FROM assessment;")
	if err != nil{
		return nil,errors.new("Could Not query the DB!")
	}
	defer rows.Close()

	a := []entity.Assessment{}

	for rows.Next(){
		assessments := entity.Assessment{}
		err = rows.Scan(&Assessment.StudentID,&Assessment.SubjectID,&Assessment.Value)
		if err != nil{
			return nil, err
		}
		a = append(a,assessments)
	}
	return a,nil
}  
