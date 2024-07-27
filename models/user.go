package models

import (
	"errors"

	"github.com/21amir21/event-booking/db"
	"github.com/21amir21/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

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

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userID

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	if err := row.Scan(&u.ID, &retrivedPassword); err != nil {
		return errors.New("credentials are invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if passwordIsValid {
		return nil
	}

	return errors.New("credentials are invalid")
}
