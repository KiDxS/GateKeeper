package models

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

func connect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./internal/storage/database.db")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return
}
