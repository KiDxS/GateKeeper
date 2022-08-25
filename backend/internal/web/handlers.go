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

	type user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	u := &user{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		serveInteralServerError(w, err)
	}
	log.Info().Msgf("%q", u)
	username, ok := models.QueryUser(u.Username, u.Password)
	if ok {
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

		w.WriteHeader(http.StatusOK)
		log.Info().Msg(jwtToken)
		return
	}

}

func handleChangePassword(w http.ResponseWriter, r *http.Request) {
	type passwordChange struct {
		CurrentPassword    string `json:"current_password"`
		NewPassword        string `json:"new_password"`
		ConfirmNewPassword string `json:"confirm_password"`
	}
	p := passwordChange{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		serveInteralServerError(w, err)
	}
	sendJSONResponse(w, 200, true, "The password has been changed.", nil)
}
