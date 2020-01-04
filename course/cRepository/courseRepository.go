package cRepository


import (
	"database/sql"
	"errors"

	"github.com/robi_a21/Cassiopeia/entity"
)


type PsqlCourseRepositoryImpl struct {
	conn *sql.DB
}

func NewPsqlCourseRepositoryImpl(Conn *sql.DB) *PsqlCourseRepositoryImpl {
	return &PsqlCourseRepositoryImpl{conn: Conn}
}


func (pr *PsqlCourseRepositoryImpl) AddCourse(course entity.Course) error {

	_, err := pr.conn.Exec("insert into course (coursename,courseid) values($1, $2)", course.CourseName, course.CourseID)
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
		err = rows.Scan(&course.CourseName,&course.CourseID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses,err
}
