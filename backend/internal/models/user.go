package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

// CheckHashedPassword is a function that compares the hashed password from the database with its possible plain text password.
func CheckHashedPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// QueryUser is a function that is used to query the database for an user. This function takes two arguments which are "username" and "password".
func (user *User) QueryUser(username, password string) (string, bool) {
	db := connect()

	stm, _ := db.Prepare("SELECT id, username, password FROM user WHERE username = ?")

	err := stm.QueryRow(username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return "", false
	}
	if !CheckHashedPassword(user.Password, password) {
		return "", false
	}
	return user.Username, true
}

// ChangeUserPassword is a function that is used to change the user's password in the database. This function takes three arguments which are "username", "currentPassword", and "newPassword".
func (user *User) ChangeUserPassword(currentPassword, newPassword string) (bool, error) {
	db := connect()

	// Checks if the password inputted is equivalent to the hashed password in the database
	if !CheckHashedPassword(user.Password, currentPassword) {
		return false, nil
	}

	// Hashes the new password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 10)

	// SQL Query to update the password column, if the conditions are right.
	stm, _ := db.Prepare("UPDATE user SET password = ? where username = ?")
	result, err := stm.Exec(hashedPassword, user.Username)
	log.Info().Msgf("%s", user.Username)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected != 1 {
		return false, nil
	}
	return true, nil
}
