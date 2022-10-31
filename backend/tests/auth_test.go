package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KiDxS/GateKeeper/internal/web/handlers/auth"
)

func sendLoginRequest(field auth.LoginFields, t testing.TB) int {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(field)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", &b)
	auth.HandleLogin(w, req)
	resp := w.Result()
	return resp.StatusCode
}

func TestAuth(t *testing.T) {
	t.Run("Test if login is successful", func(t *testing.T) {
		loginFields := auth.LoginFields{Username: "admin", Password: "abcdefghijklmd"}
		got := sendLoginRequest(loginFields, t)
		want := 204
		if got != want {
			t.Errorf("Failed to login")
		}

	})
	t.Run("Test if login is unsuccesful", func(t *testing.T) {
		loginFields := auth.LoginFields{Username: "admin", Password: "abcdefghijklm"}
		got := sendLoginRequest(loginFields, t)
		want := 401
		if got != want {
			t.Errorf("An error occured.")
		}
	})
}
