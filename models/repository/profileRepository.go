package repository

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

	rows, err := pr.conn.Query("select * from student;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	students := []entity.Student{}

	for rows.Next() {
		student := entity.Student{}
		err = rows.Scan(&student.UserName, &student.Password, &student.FirstName, &student.LastName, &student.ID, &student.Email,&student.Grade, &student.Image)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, err
}


func (pr *PsqlProfileRepositoryImpl) EmailExists(email string) bool {
	row := pr.conn.QueryRow("SELECT * FROM student WHERE email = $1", email)

	student := entity.Student{}

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Image)
	if err != nil {
		panic(err)
	}


	return true
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
		err = rows.Scan(&family.FirstName, &family.LastName, &family.Password,&family.FamilyID, &family.Phone, &family.Phone, &family.Image)
		if err != nil {
			return nil, err
		}
		families = append(families, family)
	}

	return families, err
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
		err = rows.Scan(&teacher.UserName, &teacher.Password, &teacher.Phone, &teacher.Email, &teacher.FirstName, &teacher.LastName, &teacher.TeacherID, &teacher.Image)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, err
}

func (pr *PsqlProfileRepositoryImpl) Student(id int) (entity.Student, error) {

	row := pr.conn.QueryRow("SELECT * FROM student WHERE id = $1", id)

	student := entity.Student{}

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email,&student.Grade, &student.Image)
	if err != nil {
		return student, err
	}

	return student, nil
}

func (pr *PsqlProfileRepositoryImpl) Teacher(id int) (entity.Teacher, error) {

	row := pr.conn.QueryRow("SELECT * FROM teacher WHERE id = $1", id)

	teacher := entity.Teacher{}

	err := row.Scan(&teacher.UserName, &teacher.Password, &teacher.Phone,&teacher.Email, &teacher.FirstName, &teacher.LastName, &teacher.TeacherID, &teacher.Image)
	if err != nil {
		return teacher, err
	}

	return teacher, nil
}

func (pr *PsqlProfileRepositoryImpl) Family(id int) (entity.Family, error) {

	row := pr.conn.QueryRow("SELECT * FROM teacher WHERE id = $1", id)

	family := entity.Family{}
	err := row.Scan(&family.FirstName, &family.LastName, &family.Password,&family.FamilyID, &family.Phone, &family.Phone, &family.Image)
	if err != nil {

		return family, err
	}

	return family, nil
}

func (pr *PsqlProfileRepositoryImpl) Admin(id int) (entity.Admin, error) {

	row := pr.conn.QueryRow("SELECT * FROM admin WHERE id = $1", id)

	admin := entity.Admin{}

	err := row.Scan(&admin.UserName,&admin.Password,&admin.FirstName,&admin.LastName,&admin.Email,&admin.Image)
	if err != nil {
		return admin, err
	}

	return admin, nil
}


func (pr *PsqlProfileRepositoryImpl) AdminByEmail(email string) (entity.Admin, error) {

	row := pr.conn.QueryRow("SELECT * FROM admin WHERE email = $1", email)

	admin := entity.Admin{}

	err := row.Scan(&admin.UserName,&admin.Password,&admin.FirstName,&admin.LastName,&admin.Email,&admin.Image)
	if err != nil {
		return admin, err
	}

	return admin, nil
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
		err = rows.Scan(&admin.UserName, &admin.Password, &admin.FirstName, &admin.LastName, &admin.Email, &admin.Image)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, err
}

func (pr *PsqlProfileRepositoryImpl) DeleteStudent(id int) error {

	_, err := pr.conn.Exec("DELETE FROM student WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

func (pr *PsqlProfileRepositoryImpl) DeleteTeacher(id int) error {

	_, err := pr.conn.Exec("DELETE FROM teacher WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}


