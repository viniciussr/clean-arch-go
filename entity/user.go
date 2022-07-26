package entity

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewUser create a new user
func NewUser(id int64, email, password, firstName, lastName string) *User {
	return &User{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
		Password:  password,
	}
}
