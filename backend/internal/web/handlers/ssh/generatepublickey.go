package ssh

import (
	"crypto/rsa"

	"golang.org/x/crypto/ssh"
)

// Generates a public key and returns a string
func GeneratePublicKey(privateKey *rsa.PrivateKey) string {
	publicKeyStruct, _ := ssh.NewPublicKey(privateKey.Public())
	publicKey := string(ssh.MarshalAuthorizedKey(publicKeyStruct))
	return publicKey
}
