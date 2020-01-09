package entity

type Course struct {
	CourseName  string  `json:"name""`
	CourseID    int      `json:"id"`
	Grade int             `json:"grade"`
}
