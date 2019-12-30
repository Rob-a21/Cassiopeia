package repository

import (
	"database/sql"
	"errors"
	"github.com/Rob-a21/Cassiopeia/entity"
)


type PsqlRegistrationRepositoryImpl struct {
	conn *sql.DB
}

func NewPsqlRegistrationRepositoryImpl(Conn *sql.DB) *PsqlRegistrationRepositoryImpl {
	return &PsqlRegistrationRepositoryImpl{conn: Conn}
}

func (pr *PsqlRegistrationRepositoryImpl) Students() ([]entity.Student, error) {

	rows, err := pr.conn.Query("SELECT * FROM student;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	students := []entity.Student{}

	for rows.Next() {
		student := entity.Student{}
		err = rows.Scan(student.UserName,student.Password,student.FirstName,student.LastName,student.ID,student.Image,student.Gender,student.Grade,student.Phone,student.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (pr *PsqlRegistrationRepositoryImpl) Student(id int) (entity.Student, error) {

	row := pr.conn.QueryRow("SELECT * FROM student WHERE id = $1", id)

	student := entity.Student{}

	err := row.Scan(student.UserName,student.Password,student.FirstName,student.LastName,student.ID,student.Image,student.Gender,student.Grade,student.Phone,student.Email)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (pr *PsqlRegistrationRepositoryImpl) UpdateStudent(student entity.Student) error {

	_, err := pr.conn.Exec("UPDATE student SET username=$1,password=$2, firstname=$3, lastname=$4,studentid=$5, image=$6,gender=$7,grade=$8,phone=$9, email=$10 WHERE id=$4", student.UserName,student.Password,student.FirstName,student.LastName,student.ID,student.Image,student.Gender,student.Grade,student.Phone,student.Email)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

func (pr *PsqlRegistrationRepositoryImpl) DeleteStudent(id int) error {

	_, err := pr.conn.Exec("DELETE FROM student WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

func (pr *PsqlRegistrationRepositoryImpl) RegisterStudent(student entity.Student) error {

	_, err := pr.conn.Exec("insert into student (username,password,firstname,lastname,studentid,image,gender,grade,phone,email) values($1, $2, $3,$4, $5, $6, $7, $8,$9, $10)", student.UserName,student.Password,student.FirstName,student.LastName,student.ID,student.Image,student.Gender,student.Grade,student.Phone,student.Email)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
