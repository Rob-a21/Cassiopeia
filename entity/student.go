package entity

//Student struct
type Student struct{
	UserName  string	`json:"username" gorm:"type:varchar(255);not null; unique"`
	Password  string	`json:"password" gorm:"type:varchar(255);not null; unique"`
	FirstName string	`json:"firstname" gorm:"type:varchar(255)"`
	LastName  string	`json:"lastname" gorm:"type:varchar(255)"`
	ID        int		`json:"id" gorm:"type:integer;not null; unique"`
	Image     string	`json:"image" gorm:"type:varchar(255)"`
	Gender    string	`json:"gender" gorm:"type:varchar(255)"`
	Grade     string	`json:"grade" gorm:"type:varchar(255)"`
	Phone     string	`json:"phone" gorm:"type:varchar(255)"`
	Email     string	`json:"email" gorm:"type:varchar(255)"`
}
