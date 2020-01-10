package entity

import "time"

type Attendance struct {
	Date      time.Time
	StudentId string
}