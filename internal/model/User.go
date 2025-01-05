package model

import (
	"events-api/db"
	"events-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
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
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var id int64
	var password string
	err := row.Scan(&id, &password)
	if err != nil {
		return err
	}

	err = utils.ComparePasswords(password, u.Password)
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}
