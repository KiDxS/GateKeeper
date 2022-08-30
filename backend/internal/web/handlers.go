package web

import (
	"encoding/json"
	"fmt"
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

	loginFields := &LoginFields{}
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&loginFields)
	if err != nil {
		serveInteralServerError(w, err)
		return
	}
	log.Info().Msgf("%q", loginFields)
	username, ok := user.QueryUser(loginFields.Username, loginFields.Password)
	if !ok {
		sendJSONResponse(w, 401, false, "Username or password is incorrect.", nil)
	}
	jwtToken, err := generateToken(username)
	expirationTime := time.Now().Add(1 * time.Hour)
	if err != nil {
		serveInteralServerError(w, err)
		return
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
	user := models.User{}
	changePasswordFields := ChangePasswordFields{}
	err := json.NewDecoder(r.Body).Decode(&changePasswordFields)
	if err != nil {
		serveInteralServerError(w, err)
		return
	}

	validationError := Validate(changePasswordFields)

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
		// Be careful of %s formats because it doesn't escape special characters.
		username := fmt.Sprintf("%s", JWTClaims["username"])
		password := changePasswordFields.NewPassword
		err = user.ChangeUserPassword(username, password)
		if err != nil {
			serveInteralServerError(w, err)
			return
		}

	}
	sendJSONResponse(w, 200, true, "The password has been changed.", nil)
}
