package repository

import (
	"database/sql"
	"errors"
	"github.com/Rob-a21/Cassiopeia/entity"
)

// GradeRepositoryImpl implements the assessment.gradeRepository interface
type GradeRepositoryImpl struct {
	conn *sql.DB
}

// NewGradeRepositoryImpl will create an object of PsqlGradeRepository
func NewGradeRepositoryImpl(Conn *sql.DB) *GradeRepositoryImpl {
	return &GradeRepositoryImpl{conn: Conn}
}

// StoreGrade stores new assessment information to database
func (cri *GradeRepositoryImpl) StoreGrade(c entity.Assessment) error {

	_, err := cri.conn.Exec("INSERT INTO assessment (studentid,subjectid,value,assessment) values($1, $2, $3, $4)", c.StudentID, c.SubjectID, c.Value,c.Grade)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

// UpdateGrade updates a given object with a new data
func (cri *GradeRepositoryImpl) UpdateGrade(c entity.Assessment) error {

	_, err := cri.conn.Exec("UPDATE assessment SET subjectid=$1, value=$2, assessment=$3 WHERE studentid=$4", c.SubjectID, c.Value, c.Grade, c.StudentID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteGrade removes an assessment from a database by its student and subject id
func (cri *GradeRepositoryImpl) DeleteGrade(studentid int,subjectid int) error {

	_, err := cri.conn.Exec("DELETE FROM assessment WHERE studentid=$1, subjectid=$2", studentid,subjectid)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// DeleteGrades removes assessments from a database by its id
func (cri *GradeRepositoryImpl) DeleteGrades(studentid int) error {

	_, err := cri.conn.Exec("DELETE FROM assessment WHERE studentid=$1", studentid)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// Categories returns all cateogories from the database
func (cri *GradeRepositoryImpl) Assessments(grade string) ([]entity.Assessment, error) {

	rows, err := cri.conn.Query("SELECT * FROM categories;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Assessment{}

	for rows.Next() {
		result := entity.Assessment{}
		err = rows.Scan(&result.StudentID, &result.SubjectID, &result.Value, &result.Grade)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, result)
	}

	return ctgs, nil
}

// SingleStudentAssessments returns all grades from the database of a single student
func (cri *GradeRepositoryImpl) SingleStudentAssessments(studentid int) ([]entity.Assessment, error) {

	rows, err := cri.conn.Query("SELECT * FROM categories WHERE id=$1;",studentid)
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Assessment{}

	for rows.Next() {
		result := entity.Assessment{}
		err = rows.Scan(&result.StudentID, &result.SubjectID, &result.Value, &result.Grade)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, result)
	}

	return ctgs, nil
}

// Assessment returns  grades from the database of a single student with a single subject
func (cri *GradeRepositoryImpl) Assessment(studentid int, subjectid int) ([]entity.Assessment, error) {

	rows, err := cri.conn.Query("SELECT * FROM categories WHERE studentid=$1,subjectid=$2;",studentid,subjectid)
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Assessment{}

	for rows.Next() {
		result := entity.Assessment{}
		err = rows.Scan(&result.StudentID, &result.SubjectID, &result.Value, &result.Grade)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, result)
	}

	return ctgs, nil
}
