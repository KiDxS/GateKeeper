package main

import (
	"log"
	"net/http"

	"github.com/KiDxS/GateKeeper/internal/web"
)

func main() {
	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", web.Routes())
}
