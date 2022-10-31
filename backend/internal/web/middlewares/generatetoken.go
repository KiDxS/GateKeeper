package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string) (tokenString string, err error) {

	secretKey := loadSecretKey()
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &jwtClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
