package auth

import (
	"time"

	"github.com/KiDxS/GateKeeper/internal/web/helpers"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string) (tokenString string, err error) {

	secretKey := helpers.LoadSecretKey()
	expirationTime := time.Now().Add(1 * time.Hour)
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
