package user

import "github.com/Rob-a21/Cassiopeia/entity"

type RegistrationService interface {
	RegisterStudent(student entity.Student) error
	RegisterFamily(family entity.Family) error
	RegisterTeacher(teacher entity.Teacher) error
	RegisterAdmin(admin entity.Admin) error
}

type NotificationService interface {
	GetNotification() ([]entity.Notification, error)
	Notification(id string) (entity.Notification, error)
	AddNotification(entity.Notification) error
}

type ProfileService interface {
	Students() ([]entity.Student, error)
	Student(id int) (entity.Student, error)
	DeleteStudent(id int) error
	Families() ([]entity.Family, error)
	Teachers() ([]entity.Teacher, error)
	Teacher(id string) (entity.Teacher, error)
	DeleteTeacher(id string) error
	Admin(id int) (entity.Admin, error)
	AdminByEmail(email string) (*entity.Admin, []error)
	Admins() ([]entity.Admin, error)
}


type AssessmentService interface {
	Assessments(grade string) ([]entity.Assessment,error)
	SingleStudentAssessment(id int) ([]entity.Assessment, error)
	Assessment(assessment entity.Assessment) ([]entity.Assessment,error)
	UpdateGrade(assessment entity.Assessment) error
	DeleteGrade(studentID int, subjectID int) error
	DeleteGrades(studentID int) error
	StoreGrade(assessment entity.Assessment) error
	IsQualified(studentID int) (bool, error)
}
type StudentAttendanceService interface {
	ShowAttendance() ([]entity.Attendance, error) //Categories
	CheckAttendance(id int) (entity.Attendance, error) //Category
	FillAttendance(attendance entity.Attendance) error //StoreStudent
}

type CourseService interface {
	AddCourse(course entity.Course) error
	GetCourse() ([]entity.Course, error)
	Course(id int) (entity.Course, error)
	UpdateCourse(course entity.Course) error
	DeleteCourse(id int) error
}

type SessionService interface {
	Session(sessionID string) (*entity.Session, error)
	StoreSession(session *entity.Session) (*entity.Session, error)
	DeleteSession(sessionID string) (*entity.Session, error)
}