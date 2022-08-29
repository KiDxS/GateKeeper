package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KiDxS/GateKeeper/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

// Handles the /api/v1/user/login route
func handleLogin(w http.ResponseWriter, r *http.Request) {

	u := &LoginFields{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		serveInteralServerError(w, err)
	}
	log.Info().Msgf("%q", u)
	username, ok := models.QueryUser(u.Username, u.Password)
	if !ok {
		sendJSONResponse(w, 401, false, "Username or password is incorrect.", nil)
	}
	jwtToken, err := generateToken(username)
	expirationTime := time.Now().Add(1 * time.Hour)
	if err != nil {
		serveInteralServerError(w, err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    jwtToken,
		Expires:  expirationTime,
		HttpOnly: true,
		Path:     "/",
	})

	w.WriteHeader(204)
	log.Info().Msg(jwtToken)
}

// Handles the /api/v1/user/change-password route
func handleChangePassword(w http.ResponseWriter, r *http.Request) {

	p := ChangePasswordFields{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		serveInteralServerError(w, err)
	}

	validationError := Validate(p)

	if validationError != "" {
		sendJSONResponse(w, 400, false, validationError, nil)
		return
	}

	authTokenCookie, err := r.Cookie("authToken")
	if err != nil {
		serveInteralServerError(w, err)
		return
	}
	token, _ := jwt.ParseWithClaims(authTokenCookie.Value, jwt.MapClaims{}, nil)
	JWTClaims, ok := token.Claims.(jwt.MapClaims)
	if ok {

	}
	log.Info().Msgf("%q", JWTClaims["username"])
	sendJSONResponse(w, 200, true, "The password has been changed.", nil)
}
