package ssh

// Generates a SSH pair: A Private Key and a Public Key.
//
// Function returns a privateKey and a publicKey.
func GenerateSSHPair(password string) (string, string) {
	privateKey, privateKeyStruct := GeneratePrivateKey(password)
	publicKey := GeneratePublicKey(privateKeyStruct)
	return privateKey, publicKey
}
