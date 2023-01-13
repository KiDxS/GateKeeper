package ssh

import (
	"crypto/rsa"
	"encoding/pem"

	"github.com/youmark/pkcs8"
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
