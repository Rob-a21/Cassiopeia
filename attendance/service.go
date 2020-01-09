//Service

package attendance

import "github.com/Rob-a21/Cassiopeia/entity"

type StudentAttendanceService interface {
	ShowAttendance() ([]entity.Attendance, error) //Categories
	CheckAttendance(id int) (entity.Attendance, error) //Category
	FillAttendance(category entity.Student) error //StoreStudent
}
