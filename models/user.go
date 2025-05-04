package models

import (
	"errors"
	"fmt"
	"time"

	"bojana.dev/api/db"
	"bojana.dev/api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type Login struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

func SaveLogin(userID int64, token string) (*Login, error) {
	query := `INSERT INTO logins(user_id, token, created_at) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	hashedToken := utils.HashToken(token)
	result, err := stmt.Exec(userID, hashedToken, time.Now())
	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	login := &Login{
		ID:     lastInsertId,
		UserID: userID,
	}

	return login, nil
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

func GetRegisteredUsersForEvent(eventID int64) (*[]User, error) {
	query := `SELECT u.id, u.email FROM users u
	JOIN registrations r ON u.id = r.user_id
	WHERE r.event_id = ?`

	rows, err := db.DB.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
