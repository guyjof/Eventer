package models

import (
	"errors"

	"example.com/eventer-api/db"
	"example.com/eventer-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
