package models

import (
	"errors"
	"fmt"

	"bojana.dev/api/db"
	"bojana.dev/api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u User) Save() error {
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
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println("Error executing query: ", err)
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, email, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return &User{}, err
	}

	if user.ID == 0 {
		return &User{}, fmt.Errorf("user with email %s not found", email)
	}

	return &user, nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	valid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !valid {
		return errors.New("invalid credentials")
	}

	return nil
}
