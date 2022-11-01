package auth

import "net/http"

// setCookie sets a cookie.
func setCookie(w http.ResponseWriter, cookieName, cookieValue string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		MaxAge:   60 * 60,
		HttpOnly: false,
		Path:     "/",
	})
}
