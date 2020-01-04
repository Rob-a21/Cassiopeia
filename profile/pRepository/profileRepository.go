
package pRepository

import (
	"database/sql"
	"errors"

	"github.com/robi_a21/Cassiopeia/entity"
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
		err = rows.Scan(&student.UserName,&student.Password,&student.FirstName, &student.LastName,&student.ID, &student.Email,&student.Image)
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

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName,&student.Email,&student.Image)
	if err != nil {
		return student, err
	}

	return student, nil
}



func (pr *PsqlProfileRepositoryImpl) Families() ([]entity.Family, error) {

	rows, err := pr.conn.Query("select * from family;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	var families = []entity.Family{}
	for rows.Next() {
		family := entity.Family{}
		err = rows.Scan(&family.FirstName,&family.LastName,&family.Password,&family.Phone,&family.Phone,&family.Image)
		if err != nil {
			return nil, err
		}
		families = append(families, family)
	}

	return families,err
}

func (pr *PsqlProfileRepositoryImpl) Teachers() ([]entity.Teacher, error) {

	rows, err := pr.conn.Query("select * from teacher;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	teachers := []entity.Teacher{}

	for rows.Next() {
		teacher := entity.Teacher{}
		err = rows.Scan(&teacher.UserName,&teacher.Password,&teacher.Phone,&teacher.FirstName,&teacher.LastName,&teacher.TeacherID,&teacher.Image)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers,err
}
func (pr *PsqlProfileRepositoryImpl) Admins() ([]entity.Admin, error) {

	rows, err := pr.conn.Query("select * from admin;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	admins := []entity.Admin{}

	for rows.Next() {
		admin := entity.Admin{}
		err = rows.Scan(&admin.UserName,&admin.Password,&admin.FirstName,&admin.LastName,&admin.Email,&admin.Image)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins,err
}

