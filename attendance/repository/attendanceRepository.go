package repository

import (
	"database/sql"
	"errors"
	"github.com/Rob-a21/Cassiopeia/entity"
	"time"
)

type StudentAttendanceRepositoryImpl struct {
	conn *sql.DB
}

func NewStudentAttendanceRepositoryImpl(Conn *sql.DB) *StudentAttendanceRepositoryImpl {
	return &StudentAttendanceRepositoryImpl{conn: Conn}
}

func (att *StudentAttendanceRepositoryImpl) ShowAttendance() ([]entity.Attendance, error) {

	rows, err := att.conn.Query("SELECT * FROM attendance;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	attendances := []entity.Attendance{}

	for rows.Next() {
		attendance := entity.Attendance{}
		err = rows.Scan(&attendance.StudentId, &attendance.Date)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func (att *StudentAttendanceRepositoryImpl) CheckAttendance(id int) (entity.Attendance, error) {

	row := att.conn.QueryRow("SELECT * FROM categories WHERE id = $1", id)

	c := entity.Attendance{}

	err := row.Scan(&c.StudentId, &c.Date)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (att *StudentAttendanceRepositoryImpl) FillAttendance(student entity.Student) error {

	_, err := att.conn.Exec("INSERT INTO attendance (date,id) values($1, $2)", time.Now(),student.ID)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
