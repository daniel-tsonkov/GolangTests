package models

import (
	"errors"

	"example.com/m/v2/db"
	"example.com/m/v2/utils"
)

type User struct {
	ID       int64
	Email    string `dinding: "required"`
	Password string `dinding: "required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivePasswor string
	err := row.Scan(&u.ID, &retrivePasswor)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passworIsValid := utils.CheckPasswordHash(u.Password, retrivePasswor)

	if !passworIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
