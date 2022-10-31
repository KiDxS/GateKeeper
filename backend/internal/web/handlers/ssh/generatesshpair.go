package ssh

// GenerateSSHPair is used to generate a SSH pair. It takes in a password as its argument to construct the private key.
func GenerateSSHPair(password string) (string, string) {
	privateKey, privateKeyStruct := GeneratePrivateKey(password)
	publicKey := GeneratePublicKey(privateKeyStruct)
	return privateKey, publicKey
}
