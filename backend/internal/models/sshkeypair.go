package models

type SSHKeyPair struct {
	ID      int
	label   string
	pubKey  string
	privKey string
}
