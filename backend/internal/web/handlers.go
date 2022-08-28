package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KiDxS/GateKeeper/internal/models"
	"github.com/rs/zerolog/log"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

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

func handleChangePassword(w http.ResponseWriter, r *http.Request) {

	p := ChangePasswordFields{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		serveInteralServerError(w, err)
	}
	// validationError := validateChangePasswordFields(w, &p)
	// if validationError != nil {
	// 	sendJSONResponse(w, 400, false, validationError.Error(), nil)
	// }
	validationError := Validate(p)

	if validationError != "" {
		sendJSONResponse(w, 400, false, validationError, nil)
		return
	}
	sendJSONResponse(w, 200, true, "The password has been changed", nil)
}
