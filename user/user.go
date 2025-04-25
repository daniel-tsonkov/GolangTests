package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDay  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDay)
}

func (u *User) ClearUserName() {
	u.firstName = "N/A"
	u.lastName = "N/A"
	u.birthDay = "N/A"
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "N/A",
			lastName:  "N/A",
			birthDay:  "N/A",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Fill")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDay:  birthdate,
		createdAt: time.Now(),
	}, nil
}
