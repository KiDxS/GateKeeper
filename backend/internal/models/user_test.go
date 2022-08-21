package models

import "testing"

func TestUser(t *testing.T) {
	_, got := QueryUser("admin", "admin")
	want := true

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
