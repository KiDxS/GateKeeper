package helpers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func LoadSecretKey() []byte {
	var path string
	pwd, _ := os.Getwd()
	basePath := filepath.Dir(pwd)
	if strings.Contains(basePath, "backend") {
		path = filepath.Join(basePath, ".env")
	} else {
		path = filepath.Join(basePath, "./backend/.env")
	}

	err := godotenv.Load(path)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	key := os.Getenv("SECRET_KEY")
	secretKey := []byte(key)
	return secretKey
}
