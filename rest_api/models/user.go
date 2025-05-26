package models

import "example.com/m/v2/db"

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

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}
