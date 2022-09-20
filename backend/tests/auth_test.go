package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KiDxS/GateKeeper/internal/web"
)

func sendLoginRequest(field web.LoginFields, t testing.TB) string {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(field)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", &b)
	web.HandleLogin(w, req)
	resp := w.Result()
	return resp.Status
}

func TestAuth(t *testing.T) {
	t.Run("Test if login is successful", func(t *testing.T) {
		loginFields := web.LoginFields{Username: "admin", Password: "abcdefghijklm"}
		got := sendLoginRequest(loginFields, t)
		want := "204 No Content"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

	})
	t.Run("Test if login is unsuccesful", func(t *testing.T) {
		loginFields := web.LoginFields{Username: "admin", Password: "abcdefghijklmd"}
		got := sendLoginRequest(loginFields, t)
		want := "401 Unauthorized"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
