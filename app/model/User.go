package model

type User struct {
	UserID    int
	FirstName string
	LastName  string
	Email     string
	Address   string
	Phone     string
	Password  []byte
}
