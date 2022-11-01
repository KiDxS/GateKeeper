package main

import (
	"net/http"
	"time"

	"github.com/KiDxS/GateKeeper/internal/web"
	"github.com/rs/zerolog/log"
)

func main() {
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           web.Routes(),
	}

	err := server.ListenAndServe()
	log.Info().Msg("Server is listening on port 8080")
	if err != nil {
		log.Panic().Msg(err.Error())
	}
}
