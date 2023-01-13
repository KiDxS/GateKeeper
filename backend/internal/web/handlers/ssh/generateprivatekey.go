package ssh

import (
	"crypto/rand"
	"crypto/rsa"
)

// Generates a private key
//
// Function returns a string: privateKey and a struct.

func GeneratePrivateKey(password string) (string, *rsa.PrivateKey) {
	bitSize := 2048
	privateKeyStruct, _ := rsa.GenerateKey(rand.Reader, bitSize)
	privateKey := convertPrivateKey(privateKeyStruct, password)
	return privateKey, privateKeyStruct
}
