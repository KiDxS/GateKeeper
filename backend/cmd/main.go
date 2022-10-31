package main

import (
	"log"
	"net/http"
	"time"

	"github.com/KiDxS/GateKeeper/internal/web"
)

func main() {
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           web.Routes(),
	}
	log.Println("Server is listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
