package service

import "github.com/robi_a21/Cassiopeia/entity"

import "errors"

//StudentsAttendance creates a new attendance
type StudentsAttendance map[int]*entity.Student

//NewAttendanceCache function
func NewAttendanceCache() StudentsAttendance {
	return make(map[int]*entity.Student)
}

//ShowAttendance for students
func (s StudentsAttendance) ShowAttendance(id int) (*entity.Student, error) {
	if att, ok := s[id]; ok {
		return att, nil
	}
	return nil, errors.New("Attendance Empty")
}

//FillAttendance filling function for student
func (s StudentsAttendance) FillAttendance(student *entity.Student) error {
	if _, ok := s[student.StudentID]; !ok {
		s[student.StudentID] = student
		return nil
	}
	return errors.New("Student Filled The Attendance")
}
