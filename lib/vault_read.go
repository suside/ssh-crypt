package lib

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ssh-vault/crypto/aead"
)

// DecryptVaultWithKey decrypts vault with private id_rsa key read from path
func (v *Vault) DecryptVaultWithKey(vaultPath string, keyPath string) error {
	var encryptedData []byte
	var err error
	var vsBase interface{}
	if err = v.readPrivateKey(keyPath); err != nil {
		return err
	}
	if _, err = os.Stat(vaultPath); os.IsNotExist(err) {
		v.Plaintext = []byte("")
		return nil
	}
	if encryptedData, err = ioutil.ReadFile(vaultPath); err != nil {
		return err
	}
	vsBase, err = fromBase64(string(encryptedData))
	if err != nil {
		return fmt.Errorf("%s content does not look like base64", vaultPath)
	}
	vs, ok := vsBase.(vaultSecured)
	if !ok {
		return fmt.Errorf("%s does not look like a vault file", vaultPath)
	}
	sessionKey, _ := rsa.DecryptOAEP(
		sha1.New(),
		rand.Reader,
		v.privateKey,
		vs.EncryptedSessionKeys[crc32sum(v.privateKey.Public())],
		[]byte(""),
	)
	if sessionKey != nil {
		v.Plaintext, _ = aead.Decrypt(sessionKey, vs.EncryptedData, []byte(""))
		return nil
	}
	return fmt.Errorf("Unable to read vault %s with key %s", vaultPath, keyPath)
}

// ReadStdIn content from stdin?
func (v *Vault) ReadStdIn() {
	v.Plaintext, _ = ioutil.ReadAll(os.Stdin)
}
