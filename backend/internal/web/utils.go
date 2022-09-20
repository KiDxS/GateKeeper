package web

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"

	"github.com/youmark/pkcs8"
	"golang.org/x/crypto/ssh"
)

// Converts the privateKey to pkcs8 and to PEM format.
func convertPrivateKey(key *rsa.PrivateKey, password string) string {
	var marshalPKCS8 []byte

	// Marshalls a private key into DER-encoded PKCS#8
	marshalPKCS8, _ = pkcs8.MarshalPrivateKey(key, []byte(password), pkcs8.DefaultOpts)

	// Initializes a PEM block that uses the pkcs8
	block := &pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: marshalPKCS8,
	}
	// Encodes the block into PEM format and converts it into string
	privateKey := string(pem.EncodeToMemory(block))
	return privateKey
}

// Generates a private key
func GeneratePrivateKey(password string) (string, *rsa.PrivateKey) {
	bitSize := 2048
	privateKeyStruct, _ := rsa.GenerateKey(rand.Reader, bitSize)
	privateKey := convertPrivateKey(privateKeyStruct, password)
	return privateKey, privateKeyStruct
}

// Generates a public key and returns a string
func GeneratePublicKey(privateKey *rsa.PrivateKey) string {
	publicKeyStruct, _ := ssh.NewPublicKey(privateKey.Public())
	publicKey := string(ssh.MarshalAuthorizedKey(publicKeyStruct))
	return publicKey
}

// Generates a SSH pair: A Private Key and a Public Key
func GenerateSSHPair(password string) (string, string) {
	privateKey, privateKeyStruct := GeneratePrivateKey(password)
	publicKey := GeneratePublicKey(privateKeyStruct)
	return privateKey, publicKey
}
