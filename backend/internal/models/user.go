package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type User struct {
	ID       int
	Username string
	Password string
}

// Creates a connection to the database

// Queries the database for valid user credentials
func (user *User) QueryUser(username, password string) (u string, ifExists bool) {
	db := connect()

	stm, _ := db.Prepare("SELECT * FROM user WHERE username = ? AND password = ?")

	err := stm.QueryRow(username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return "", false
	}
	return user.Username, true

}

// Changes the password of the user
func (user *User) ChangeUserPassword(username, currentPassword, password string) (passwordUpdated bool, err error) {
	user.Username = username
	user.Password = currentPassword
	db := connect()

	// SQL Query to update the password column, if the conditions are right.
	stmt, _ := db.Prepare("UPDATE user SET password = ? where username = ? AND password = ?")
	result, err := stmt.Exec(password, user.Username, user.Password)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Info().Msgf("Rows affected: %d", rowsAffected)

	if rowsAffected != 1 {
		return false, nil
	}
	return true, nil
}
