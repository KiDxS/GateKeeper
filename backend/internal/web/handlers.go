package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KiDxS/GateKeeper/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

// Handles the /api/v1/user/login route
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	fields := &LoginFields{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		serveInteralServerError(w, err)
		return
	}
	username, ok := user.QueryUser(fields.Username, fields.Password)
	if !ok {
		sendJSONResponse(w, 401, false, "Username or password is incorrect.", nil)
		return
	}
	jwtToken, err := generateToken(username)
	// expirationTime := time.Now().Add(1 * time.Hour)
	if err != nil {
		serveInteralServerError(w, err)
		return
	}

	// Sets the cookie of the user
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    jwtToken,
		MaxAge:   60 * 60,
		HttpOnly: false,
		Path:     "/",
	})

	w.WriteHeader(204)
}

// Handles the /api/v1/user/logout route
func handleLogout(w http.ResponseWriter, _ *http.Request) {
	// Expires the authToken cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    "",
		MaxAge:   0,
		HttpOnly: false,
		Path:     "/",
	})
	w.WriteHeader(204)
}

// Handles the /api/v1/user/change-password route
func handleChangePassword(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	fields := ChangePasswordFields{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		serveInteralServerError(w, err)
		return
	}

	validationError := Validate(fields)

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
		passwordChanged, err := user.ChangeUserPassword(username, fields.CurrentPassword, fields.NewPassword)
		if err != nil {
			serveInteralServerError(w, err)
			return
		}
		if !passwordChanged {
			sendJSONResponse(w, 400, false, "Your current password is not correct.", nil)
			return
		}

	}
	sendJSONResponse(w, 200, true, "The password has been changed.", nil)
}

func HandleSSHGeneration(w http.ResponseWriter, r *http.Request) {
	fields := SSHGenerationFields{}
	keypair := models.SSHKeyPair{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		serveInteralServerError(w, err)
		return
	}
	validationError := Validate(fields)
	if validationError != "" {
		sendJSONResponse(w, 400, false, validationError, nil)
		return
	}
	privateKey, publicKey := GenerateSSHPair(fields.Password)
	err = keypair.InsertSSHPairKey(fields.Label, publicKey, privateKey)
	if err != nil {
		serveInteralServerError(w, err)
	}
	sendJSONResponse(w, 200, true, "An SSH keypair has been generated", nil)
}

func HandleRetrieveSSHKeypair(w http.ResponseWriter, r *http.Request) {
	keyID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		serveInteralServerError(w, err)
		return
	}
	keypair := models.SSHKeyPair{}
	err = keypair.QuerySSHKeyPair(keyID)
	if err != nil {
		if err == models.ErrNoRows {
			sendJSONResponse(w, 404, false, "ID doesn't exist.", nil)
			return
		}
		serveInteralServerError(w, err)
		return
	}
	sendJSONResponse(w, 200, true, "The SSH keypair has been retrieved.", keypair)
}
func HandleRetrieveSSHKeypairLabels(w http.ResponseWriter, _ *http.Request) {
	keypair := models.SSHKeyPair{}
	labels, err := keypair.QuerySSHKeyPairLabels()
	if err != nil {
		if err == models.ErrNoRows {
			sendJSONResponse(w, 404, false, "No SSH keypairs haven't been created yet.", nil)
			return
		}
		serveInteralServerError(w, err)
		return
	}
	sendJSONResponse(w, 200, true, "Retrieved a list of labels of SSH keypairs", labels)
}

func HandleDeleteSSHKeypair(w http.ResponseWriter, r *http.Request) {
	keyID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		serveInteralServerError(w, err)
		return
	}
	keypair := models.SSHKeyPair{}
	err = keypair.DeleteSSHKeyPair(keyID)
	if err != nil {
		if err == models.ErrNoRows {
			sendJSONResponse(w, 404, false, "ID doesn't exist.", nil)
			return
		}
		serveInteralServerError(w, err)
		return
	}
	sendJSONResponse(w, 200, true, "The SSH Keypair has been deleted.", nil)
}
