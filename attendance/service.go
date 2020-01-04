package attendance

import "github.com/robi_a21/Cassiopeia/entity"

type AttendanceService interface {
	ShowAttendance(id int) (*entity.Student, error)
	FillAttendace(student *entity.Student) error
}
