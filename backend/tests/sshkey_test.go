package tests

import (
	"regexp"
	"testing"

	"github.com/KiDxS/GateKeeper/internal/web"
)

func checkIfPrivateKey(privateKey string) (matches bool) {
	matches, _ = regexp.MatchString("\\s*(\\bBEGIN\\b).*(PRIVATE KEY\\b)\\s*", privateKey)
	return
}

func checkIfPublicKey(publicKey string) (matches bool) {
	matches, _ = regexp.MatchString("ssh-rsa AAAA", publicKey)
	return
}

func TestGeneration(t *testing.T) {
	t.Run("Generate Private Key", func(t *testing.T) {
		privateKey, _ := web.GenerateSSHPair("test")
		got := checkIfPrivateKey(privateKey)
		want := true
		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
	t.Run("Generate Public Key", func(t *testing.T) {
		_, publicKey := web.GenerateSSHPair("test")
		got := checkIfPublicKey(publicKey)
		want := true

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

}
