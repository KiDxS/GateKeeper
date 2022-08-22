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
	return
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	type user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	u := &user{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msgf("%q", u)
	username, ok := models.QueryUser(u.Username, u.Password)
	if ok {
		jwtToken, err := generateToken(username)
		expirationTime := time.Now().Add(1 * time.Hour)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "authToken",
			Value:    jwtToken,
			Expires:  expirationTime,
			HttpOnly: true,
			Path:     "/",
		})

		w.Write([]byte("test"))
		w.WriteHeader(http.StatusOK)
		log.Info().Msg(jwtToken)
		return
	}

}
