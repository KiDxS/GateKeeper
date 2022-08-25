package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type user struct {
	ID       int
	Username string
	Password string
}

// Queries the database for valid user credentials
func QueryUser(username, password string) (u string, ifExists bool) {
	user := &user{}
	db, err := sql.Open("sqlite3", "./internal/storage/database.db")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	stm, _ := db.Prepare("SELECT * FROM user WHERE username = ? AND password = ?")

	err = stm.QueryRow(username, password).Scan(&user.ID, &user.Username, &user.Password)
	log.Info().Msg(user.Username)
	if err == sql.ErrNoRows {
		return "", false
	}
	return user.Username, true

}
