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
//Student struct for data caching
type Student struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
	ID        int
	Email     string
	Grade     int
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
	TeacherID int
	Image     string
}


//Family struct for data caching
type Family struct {
	FirstName string
	LastName  string
	Username string
	Password  string
	FamilyID        int
	Phone     string
	Email     string
	Image    string
}
type Assessment struct {
	Value     int    `json:"value"`
	SubjectID int     `json:"subjectid"`
	StudentID int      `json:"studentid"`
	Grade     int    `json:"grade"`
}


type Attendance struct {
	Date      time.Time `json:"date"`
	StudentID int        `json:"studentid"`
}

type Course struct {
	CourseName  string  `json:"name""`
	CourseID    int      `json:"id"`
	Grade int             `json:"grade"`
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


