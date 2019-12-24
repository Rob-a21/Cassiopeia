
package pRepository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
)

type PsqlProfileRepositoryImpl struct {
	conn *sql.DB
}

func NewPsqlProfileRepositoryImpl(Conn *sql.DB) *PsqlProfileRepositoryImpl {
	return &PsqlProfileRepositoryImpl{conn: Conn}
}

func (pr *PsqlProfileRepositoryImpl) Students() ([]entity.Student, error) {

	rows, err := pr.conn.Query("SELECT * FROM student;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	students := []entity.Student{}

	for rows.Next() {
		student := entity.Student{}
		err = rows.Scan(&student.UserName, &student.FirstName, &student.LastName, &student.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students,err
}

func (pr *PsqlProfileRepositoryImpl) Student(id int) (entity.Student, error) {

	row := pr.conn.QueryRow("SELECT * FROM student WHERE id = $1", id)

	student := entity.Student{}

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName,&student.Email)
	if err != nil {
		return student, err
	}

	return student, nil
}

