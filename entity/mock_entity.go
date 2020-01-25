package entity

import "time"

var CourseMock = Course{
	CourseName:"Mock 01",
	CourseID:1,
	Grade :2,
}

var AttendanceMock = Attendance{
	Date:time.Now(),
	StudentID:1,
}