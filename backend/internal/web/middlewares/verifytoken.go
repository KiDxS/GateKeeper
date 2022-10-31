package middlewares

import "github.com/golang-jwt/jwt/v4"

func verifyToken(tokenString string) (token *jwt.Token, err error) {
	secretKey := loadSecretKey()
	claims := &jwtClaim{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	return
}
