package entity


type Session struct {
	ID         uint
	UUID       string `"uuid"`
	Expires    int64  `"expires"`
	SigningKey []byte `"signinkey"`
}