package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Password string
}

// Creates a connection to the database

// Queries the database for valid user credentials
// QueryUser is a function that is used to query the database for an user. This function takes two arguments which are "username" and "password".
func (user *User) QueryUser(username, password string) (string, bool) {
	db := connect()

	stm, _ := db.Prepare("SELECT * FROM user WHERE username = ? AND password = ?")

	err := stm.QueryRow(username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return "", false
	}
	return user.Username, true

}

// ChangeUserPassword is a function that is used to change the user's password in the database. This function takes three arguments which are "username", "currentPassword", and "newPassword".
func (user *User) ChangeUserPassword(username, currentPassword, newPassword string) (bool, error) {
	user.Username = username
	user.Password = currentPassword
	db := connect()

	// SQL Query to update the password column, if the conditions are right.
	stm, _ := db.Prepare("UPDATE user SET password = ? where username = ? AND password = ?")
	result, err := stm.Exec(newPassword, user.Username, user.Password)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected != 1 {
		return false, nil
	}
	return true, nil
}
