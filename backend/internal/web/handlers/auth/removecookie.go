package auth

import "net/http"

// removeCookie removes the following cookie by overwriting its value and expiring its age.
func removeCookie(w http.ResponseWriter, cookieName string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		MaxAge:   0,
		HttpOnly: false,
		Path:     "/",
	})
}
