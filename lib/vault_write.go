package lib

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"io/ioutil"
	"os"

	"os/exec"

	"github.com/ssh-vault/crypto/aead"
)

// StoreSecuredVault encrypts Vault and stores it on vaultPath
func (v *Vault) StoreSecuredVault(vaultPath string) error {
	var err error
	v.sessionKey = make([]byte, 32)
	rand.Read(v.sessionKey)
	vs := vaultSecured{EncryptedSessionKeys: make(map[uint32][]byte)}
	for _, key := range v.publicKeys {
		sessionKeySecure, _ := rsa.EncryptOAEP(
			sha1.New(),
			rand.Reader,
			key,
			v.sessionKey,
			[]byte(""),
		)
		vs.EncryptedSessionKeys[crc32sum(&key)] = sessionKeySecure
	}
	vs.EncryptedData, err = aead.Encrypt(v.sessionKey, v.Plaintext, []byte(""))
	if err != nil {
		return err
	}
	ioutil.WriteFile(vaultPath, []byte(toBase64(vs)), 0644)
	return nil
}

// EditVaultFile vault.Path file
func (v *Vault) EditVaultFile() {
	tmpfile, _ := ioutil.TempFile("", "ssh-crypt")
	defer os.Remove(tmpfile.Name())
	tmpfile.Write([]byte(v.Plaintext))
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}
	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	v.Plaintext, _ = ioutil.ReadFile(tmpfile.Name())
}
