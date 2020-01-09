//Repository

package attendance

import "github.com/Rob-a21/Cassiopeia/entity"

type StudentAttendanceRepository interface {
	ShowAttendance() ([]entity.Attendance, error) //Categories
	CheckAttendance(id int) (entity.Attendance, error) //Category
	FillAttendance(category entity.Student) error //StoreStudent
}



