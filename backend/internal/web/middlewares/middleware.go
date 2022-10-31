package middlewares

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type jwtClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func loadSecretKey() []byte {
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
