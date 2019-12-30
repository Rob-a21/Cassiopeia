package entity

//Admin struct
type Admin struct {
	UserName  string `json:"admin-name"`
	Password  int	 `json:"password" gorm:"type:varchar(255)"`
	FirstName string `json:"firstname" gorm:"type:varchar(255)"`
	LastName  string `json:"lastname" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(255);not null; unique"`
}
