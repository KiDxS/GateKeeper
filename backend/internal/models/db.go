package models

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

var ErrNoRows error

func connect() *sql.DB {
	var db *sql.DB
	var path string
	pwd, _ := os.Getwd()
	basePath := filepath.Dir(pwd)
	if strings.Contains(basePath, "backend") {
		path = filepath.Join(basePath, "./internal/storage/database.db")
		db, err := sql.Open("sqlite3", path)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		return db
	}
	path = filepath.Join(basePath, "./backend/internal/storage/database.db")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	// db, err := sql.Open("sqlite3", path)
	// if err != nil {
	// 	log.Fatal().Msg(err.Error())
	// }
	return db
}
