package models

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func connect() *sql.DB {
	pwd, _ := os.Getwd()
	rootPath := filepath.Dir(pwd)
	path := filepath.Join(rootPath, "/internal/storage/database.db")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return db
}
