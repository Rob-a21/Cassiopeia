package cRepository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
)

type PsqlCourseRepositoryImpl struct {
	conn *sql.DB
}

func NewPsqlCourseRepositoryImpl(Conn *sql.DB) *PsqlCourseRepositoryImpl {
	return &PsqlCourseRepositoryImpl{conn: Conn}
}

func (pr *PsqlCourseRepositoryImpl) AddCourse(course entity.Course) error {

	_, err := pr.conn.Exec("insert into course (coursename,courseid,grade) values($1, $2,$3)", course.CourseName, course.CourseID,course.Grade)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

func (pr *PsqlCourseRepositoryImpl) GetCourse() ([]entity.Course, error) {

	rows, err := pr.conn.Query("SELECT * FROM course;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	courses := []entity.Course{}

	for rows.Next() {
		course := entity.Course{}
		err = rows.Scan(&course.CourseName, &course.CourseID,&course.Grade)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, err
}

func (pr *PsqlCourseRepositoryImpl) Course(id int) (entity.Course, error) {

	row := pr.conn.QueryRow("SELECT * FROM course WHERE id = $1", id)

	c := entity.Course{}

	err := row.Scan(&c.CourseName, &c.CourseID)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (pr *PsqlCourseRepositoryImpl) UpdateCourse(c entity.Course) error {

	_, err := pr.conn.Exec("UPDATE course SET coursename=$1,courseid=$2, grade=$3 WHERE id=$4", c.CourseName, c.CourseID,c.Grade)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

func (pr *PsqlCourseRepositoryImpl) DeleteCourse(id int) error {

	_, err := pr.conn.Exec("DELETE FROM course WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}
