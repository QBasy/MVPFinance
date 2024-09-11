package models

import (
	"MVPFinanceApp/db"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

func CreateUser(user User) (int, error) {
	var userID int
	sqlStatement := `INSERT INTO users (name, email, password_hash, created_at) 
	VALUES ($1, $2, $3, NOW()) RETURNING id`
	err := db.DB.QueryRow(sqlStatement, user.Name, user.Email, user.PasswordHash).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func GetUserByID(id int) (*User, error) {
	var user User
	sqlStatement := `SELECT id, name, email, password_hash, created_at FROM users WHERE id = $1`
	row := db.DB.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user User) error {
	sqlStatement := `UPDATE users SET name = $1, email = $2, password_hash = $3 
	WHERE id = $4`
	_, err := db.DB.Exec(sqlStatement, user.Name, user.Email, user.PasswordHash, user.ID)
	return err
}

func DeleteUser(id int) error {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := db.DB.Exec(sqlStatement, id)
	return err
}
