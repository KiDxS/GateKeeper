package ssh

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KiDxS/GateKeeper/internal/models"
	"github.com/KiDxS/GateKeeper/internal/web/helpers"
	"github.com/go-chi/chi"
)

func HandleSSHGeneration(w http.ResponseWriter, r *http.Request) {
	fields := SSHGenerationFields{}
	keypair := models.SSHKeyPair{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		helpers.ServeInteralServerError(w, err)
		return
	}
	validationError := helpers.Validate(fields)
	if validationError != "" {
		helpers.SendJSONResponse(w, 400, false, validationError, nil)
		return
	}
	privateKey, publicKey := GenerateSSHPair(fields.Password)
	err = keypair.InsertSSHPairKey(fields.Label, publicKey, privateKey)
	if err != nil {
		helpers.ServeInteralServerError(w, err)
	}
	helpers.SendJSONResponse(w, 200, true, "An SSH keypair has been generated", nil)
}

func HandleRetrieveSSHKeypair(w http.ResponseWriter, r *http.Request) {
	keyID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServeInteralServerError(w, err)
		return
	}
	keypair := models.SSHKeyPair{}
	err = keypair.QuerySSHKeyPair(keyID)
	if err != nil {
		if err == models.ErrNoRows {
			helpers.SendJSONResponse(w, 404, false, "ID doesn't exist.", nil)
			return
		}
		helpers.ServeInteralServerError(w, err)
		return
	}
	helpers.SendJSONResponse(w, 200, true, "The SSH keypair has been retrieved.", keypair)
}
func HandleRetrieveSSHKeypairLabels(w http.ResponseWriter, _ *http.Request) {
	keypair := models.SSHKeyPair{}
	labels, err := keypair.QuerySSHKeyPairLabels()
	if err != nil {
		if err == models.ErrNoRows {
			helpers.SendJSONResponse(w, 404, false, "No SSH keypairs haven't been created yet.", nil)
			return
		}
		helpers.ServeInteralServerError(w, err)
		return
	}
	helpers.SendJSONResponse(w, 200, true, "Retrieved a list of labels of SSH keypairs", labels)
}

func HandleDeleteSSHKeypair(w http.ResponseWriter, r *http.Request) {
	keyID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServeInteralServerError(w, err)
		return
	}
	keypair := models.SSHKeyPair{}
	err = keypair.DeleteSSHKeyPair(keyID)
	if err != nil {
		if err == models.ErrNoRows {
			helpers.SendJSONResponse(w, 404, false, "ID doesn't exist.", nil)
			return
		}
		helpers.ServeInteralServerError(w, err)
		return
	}
	helpers.SendJSONResponse(w, 200, true, "The SSH Keypair has been deleted.", nil)
}
