package repository

import (
	"database/sql"
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type AttendanceRepoMock struct {
	conn *sql.DB
}

func NewAttendanceRepoMock(db *sql.DB) models.StudentAttendanceRepository {
	return &AttendanceRepoMock{conn: db}
}

func (aRepo *AttendanceRepoMock) ShowAttendance() ([]entity.Attendance, error) {
	posts := []entity.Attendance{entity.AttendanceMock}

	return posts, nil
}

func (aRepo *AttendanceRepoMock) CheckAttendance(id int) (entity.Attendance, error) {
	null := entity.Attendance{}

	if id == 0001 {
		return entity.AttendanceMock, nil
	}

	return null, nil
}

func (aRepo *AttendanceRepoMock) FillAttendance(attendance entity.Attendance) error {
	attendance = entity.AttendanceMock

	return nil
}





