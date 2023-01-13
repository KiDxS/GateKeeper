package middlewares

import "github.com/golang-jwt/jwt/v4"

type jwtClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
