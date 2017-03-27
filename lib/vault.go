package lib

import "crypto/rsa"

// Vault struct
type Vault struct {
	Plaintext  []byte
	sessionKey []byte
	publicKeys []*rsa.PublicKey
	privateKey *rsa.PrivateKey
}

type vaultSecured struct {
	EncryptedData        []byte
	EncryptedSessionKeys map[uint32][]byte
}
