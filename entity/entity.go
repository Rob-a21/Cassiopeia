package entity

import "time"

type Admin struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
	Email     string
	Image     string

}

type Assessment struct {
	Value     int
	SubjectID string
	StudentID string
	Grade     string
}


type Attendance struct {
	Date      time.Time
	StudentId int
}

type Course struct {
	CourseName  string  `json:"name""`
	CourseID    int      `json:"id"`
	Grade int             `json:"grade"`
}


//Family struct for data caching
type Family struct {
	FirstName string
	LastName  string
	Username string
	Password  string
	Phone     string
	Email     string
	Image    string
}



type Notification struct{

	Message string             `json:"message""`

	NotifyName string           `json:"name""`

	NotificationDate time.Time   `json:"date""`
}



type Session struct {
	ID         uint
	UUID       string `"uuid"`
	Expires    int64  `"expires"`
	SigningKey []byte `"signinkey"`
}

//Student struct for data caching
type Student struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
	ID        int
	Email     string
	Image     string
}



//Teacher struct for data caching
type Teacher struct {
	UserName  string
	Password  string
	Phone     string
	Email     string
	FirstName string
	LastName  string
	TeacherID string
	Image     string
}
