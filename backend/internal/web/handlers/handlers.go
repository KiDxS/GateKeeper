package handlers

import (
	"net/http"

	authfeature "github.com/KiDxS/GateKeeper/internal/web/handlers/auth"
	sshfeature "github.com/KiDxS/GateKeeper/internal/web/handlers/ssh"
)

func RetrieveSSHKeypair(w http.ResponseWriter, r *http.Request) {
	sshfeature.HandleRetrieveSSHKeypair(w, r)
}

func SSHGeneration(w http.ResponseWriter, r *http.Request) {
	sshfeature.HandleSSHGeneration(w, r)
}

func RetrieveSSHKeypairLabels(w http.ResponseWriter, r *http.Request) {
	sshfeature.HandleRetrieveSSHKeypairLabels(w, r)
}

func DeleteSSHKeypair(w http.ResponseWriter, r *http.Request) {
	sshfeature.HandleDeleteSSHKeypair(w, r)
}

func Login(w http.ResponseWriter, r *http.Request) {
	authfeature.HandleLogin(w, r)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	authfeature.HandleChangePassword(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authfeature.HandleLogout(w, r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	authfeature.HandleIndex(w, r)
}
