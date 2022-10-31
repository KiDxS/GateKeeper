package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string) (tokenString string, err error) {

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
