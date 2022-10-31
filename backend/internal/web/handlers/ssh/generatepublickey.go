package ssh

import (
	"crypto/rsa"

	"golang.org/x/crypto/ssh"
)

// GeneratePublicKey generates a public key from a private key and outputs the result as a string.
func GeneratePublicKey(privateKey *rsa.PrivateKey) string {
	publicKeyStruct, _ := ssh.NewPublicKey(privateKey.Public())
	publicKey := string(ssh.MarshalAuthorizedKey(publicKeyStruct))
	return publicKey
}
