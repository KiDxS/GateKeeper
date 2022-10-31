package middlewares

import (
	"net/http"

	"github.com/KiDxS/GateKeeper/internal/web/helpers"
)

func AuthMiddleware(next http.Handler) (handler http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken, err := r.Cookie("authToken")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, _ := verifyToken(authToken.Value)
		if !token.Valid {
			helpers.ServeForbiddenError(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
