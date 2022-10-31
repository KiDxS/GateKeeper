package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/KiDxS/GateKeeper/internal/models"
	"github.com/KiDxS/GateKeeper/internal/web/handlers/ssh"
)

func checkIfPrivateKey(privateKey string) (matches bool) {
	matches, _ = regexp.MatchString("\\s*(\\bBEGIN\\b).*(PRIVATE KEY\\b)\\s*", privateKey)
	return
}

func checkIfPublicKey(publicKey string) (matches bool) {
	matches, _ = regexp.MatchString("ssh-rsa AAAA", publicKey)
	return
}

// Unit tests for generating the SSH keypairs
func TestSSHGeneration(t *testing.T) {
	t.Run("Generate Private Key", func(t *testing.T) {
		privateKey, _ := ssh.GenerateSSHPair("test")
		got := checkIfPrivateKey(privateKey)
		want := true
		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
	t.Run("Generate Public Key", func(t *testing.T) {
		_, publicKey := ssh.GenerateSSHPair("test")
		got := checkIfPublicKey(publicKey)
		want := true

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

}

// Unit tests for the interaction in the database
func TestSSHDataModel(t *testing.T) {
	t.Run("Insert generated SSH key pair to the database", func(t *testing.T) {
		keypair := models.SSHKeyPair{}
		err := keypair.InsertSSHPairKey("test label shhh", "test pub key shh", "test priv key shh")
		if err != nil {
			t.Errorf("An error has occured while inserting a key pair")
		}
	})
	t.Run("Retrieve the information of a single SSH pair", func(t *testing.T) {
		keypair := models.SSHKeyPair{}
		err := keypair.QuerySSHKeyPair(9)
		if err != nil {
			t.Errorf("An error has occured while fetching a keypair")
		}
	})
	t.Run("Retrieve the labels of all the SSH pairs", func(t *testing.T) {
		keypair := models.SSHKeyPair{}
		labels, _ := keypair.QuerySSHKeyPairLabels()
		if len(labels) == 0 {
			t.Errorf("We weren't able to retrieve the labels from the database")
		}
	})
	t.Run("Delete a SSH keypair", func(t *testing.T) {
		keypair := models.SSHKeyPair{}
		err := keypair.DeleteSSHKeyPair(10)
		if err != nil {
			t.Errorf("we weren't able to delete a SSH pair key from the database")
		}
	})
	t.Run("Update the label of a SSH keypair", func(t *testing.T) {
		keypair := models.SSHKeyPair{}
		err := keypair.UpdateSSHKeyPairLabel(9, "testupdates")
		if err != nil {
			t.Errorf("An error has occured while updating a label of a SSH key pair")
		}
		got := keypair.Label
		want := "testupdates"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

	})
}

func TestSSHRoute(t *testing.T) {
	t.Run("Test SSH route for successful request", func(t *testing.T) {
		fields := ssh.SSHGenerationFields{Label: "sshgen", Password: "test"}
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(fields)
		if err != nil {
			t.Fatal(nil)
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/v1/key", &b)
		ssh.HandleSSHGeneration(w, r)
		resp := w.Result()
		got := resp.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})
	t.Run("Test SSH route for a validation error", func(t *testing.T) {
		fields := ssh.SSHGenerationFields{Label: "", Password: "test"}
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(fields)
		if err != nil {
			t.Fatal(nil)
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/v1/key", &b)
		ssh.HandleSSHGeneration(w, r)
		resp := w.Result()
		got := resp.StatusCode
		want := 400

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}
