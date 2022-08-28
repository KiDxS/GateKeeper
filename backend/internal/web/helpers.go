package web

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func serveInteralServerError(w http.ResponseWriter, err error) {
	log.Fatal().Stack().Msg(err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func serveForbiddenError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func serveError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func serveNotFoundError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// Function that sends a JSON Response
func sendJSONResponse(w http.ResponseWriter, statusCode int, status bool, message string, data interface{}) {
	type JSONResponse struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	resp := JSONResponse{
		Success: status,
		Message: message,
		Data:    data,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		serveInteralServerError(w, err)
	}
	w.WriteHeader(statusCode)
	w.Write(jsonResp)
	return
}
