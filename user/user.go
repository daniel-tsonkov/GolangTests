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

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDay)
}

func (u *User) ClearUserName() {
	u.firstName = "N/A"
	u.lastName = "N/A"
	u.birthDay = "N/A"
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
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
