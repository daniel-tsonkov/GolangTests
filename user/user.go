package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDay  string
	CreatedAt time.Time
}

func (u *User) outputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDay)
}

func (u *User) clearUserName() {
	u.FirstName = "N/A"
	u.LastName = "N/A"
	u.BirthDay = "N/A"
}

func newUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Fill")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDay:  birthdate,
		CreatedAt: time.Now(),
	}, nil
}
