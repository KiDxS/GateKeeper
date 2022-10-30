package web

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func serveInteralServerError(w http.ResponseWriter, err error) {
	log.Warn().Stack().Msg(err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func serveForbiddenError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)

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
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		serveInteralServerError(w, err)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(jsonResp)
}
