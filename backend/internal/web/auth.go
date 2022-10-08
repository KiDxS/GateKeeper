package web

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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

// Verifies the JWT Token
func verifyToken(tokenString string) (token *jwt.Token, err error) {
	secretKey := loadSecretKey()
	claims := &jwtClaim{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	return
}

func authMiddleware(next http.Handler) (handler http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken, err := r.Cookie("authToken")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, _ := verifyToken(authToken.Value)
		if !token.Valid {
			serveForbiddenError(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Generates a JWT token for authentication
func generateToken(username string) (tokenString string, err error) {

	secretKey := loadSecretKey()
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwtClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
