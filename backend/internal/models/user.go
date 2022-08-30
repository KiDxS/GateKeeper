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
	// db, err := sql.Open("sqlite3", "./internal/storage/database.db")
	// if err != nil {
	// 	log.Fatal().Msg(err.Error())
	// }
	stm, _ := db.Prepare("SELECT * FROM user WHERE username = ? AND password = ?")

	err := stm.QueryRow(username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return "", false
	}
	return user.Username, true

}

func (user *User) ChangeUserPassword(username, password string) (err error) {
	user.Username = username
	user.Password = password
	db := connect()
	stmt, _ := db.Prepare("UPDATE user SET password = ? where username = ?")
	result, err := stmt.Exec(user.Password, user.Username)
	if err != nil {
		return
	}
	rowsAffected, _ := result.RowsAffected()
	log.Info().Msgf("Rows affected: %d", rowsAffected)

	// stm, _ := db.Prepare("SELECT * FROM user WHERE username = ? AND password = ?")

	// err = stm.QueryRow(username, password).Scan(&user.ID, &user.Username, &user.Password)
	// if err == sql.ErrNoRows {
	// 	return
	// }
	// log.Info().Msgf("Username: %q, Password: %q", user.Username, user.Password)

	return nil
}
