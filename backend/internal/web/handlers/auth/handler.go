package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KiDxS/GateKeeper/internal/models"
	"github.com/KiDxS/GateKeeper/internal/web/helpers"
	"github.com/KiDxS/GateKeeper/internal/web/middlewares"
	"github.com/golang-jwt/jwt/v4"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

// HandleLogin is the logic handler for the /api/v1/user/login route
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	fields := LoginFields{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		helpers.ServeInteralServerError(w, err)
		return
	}
	username, ok := user.QueryUser(fields.Username, fields.Password)
	if !ok {
		helpers.SendJSONResponse(w, 401, false, "Username or password is incorrect.", nil)
		return
	}
	jwtToken, err := middlewares.GenerateToken(username)
	// expirationTime := time.Now().Add(1 * time.Hour)
	if err != nil {
		helpers.ServeInteralServerError(w, err)
		return
	}

	// Sets the cookie of the user
	setCookie(w, "authToken", jwtToken)
	w.WriteHeader(204)
}

// HandleLogout is the logic handler for the /api/v1/user/logout route
func HandleLogout(w http.ResponseWriter, _ *http.Request) {
	// Expires the authToken cookie
	removeCookie(w, "authToken")
	w.WriteHeader(204)
}

// HandleChangePassword is the logic handler for the /api/v1/user/change-password route
func HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	fields := ChangePasswordFields{}
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

	authTokenCookie, err := r.Cookie("authToken")
	if err != nil {
		helpers.ServeInteralServerError(w, err)
		return
	}
	token, _ := jwt.ParseWithClaims(authTokenCookie.Value, jwt.MapClaims{}, nil)
	JWTClaims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		// Be careful of %s formats because it doesn't escape special characters.
		username := fmt.Sprintf("%s", JWTClaims["username"])
		passwordChanged, err := user.ChangeUserPassword(username, fields.CurrentPassword, fields.NewPassword)
		if err != nil {
			helpers.ServeInteralServerError(w, err)
			return
		}
		if !passwordChanged {
			helpers.SendJSONResponse(w, 400, false, "Your current password is not correct.", nil)
			return
		}

	}
	helpers.SendJSONResponse(w, 200, true, "The password has been changed.", nil)
}
