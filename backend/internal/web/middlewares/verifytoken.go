package middlewares

import (
	"github.com/KiDxS/GateKeeper/internal/web/helpers"
	"github.com/golang-jwt/jwt/v4"
)

func verifyToken(tokenString string) (token *jwt.Token, err error) {
	secretKey := helpers.LoadSecretKey()
	claims := &jwtClaim{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	return
}
