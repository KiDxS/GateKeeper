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
	log.Info().Msg(user.Username)
	if err == sql.ErrNoRows {
		return "", false
	}
	return user.Username, true

}

func (user *User) ChangeUserPassword(username, password string) (err error) {
	db := connect()
	stm, _ := db.Prepare("UPDATE user SET password = ? where username = ?")
	_, err = stm.Exec(password, username)
	if err != nil {
		return
	}

	return nil
}
