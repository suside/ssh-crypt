package lib

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"strings"

	ssh "github.com/ianmcmahon/encoding_ssh"
)

// ReadAuthorizedKeys file from path
func (v *Vault) ReadAuthorizedKeys(path string) {
	authorizedKeys, _ := ioutil.ReadFile(path)
	authorizedKeysList := strings.Split(strings.TrimSpace(string(authorizedKeys)), "\n")
	for _, authorizedKey := range authorizedKeysList {
		if pubKey, err := ssh.DecodePublicKey(authorizedKey); err == nil {
			// TODO handle other key types
			v.publicKeys = append(v.publicKeys, pubKey.(*rsa.PublicKey))
		}
	}
}

func (v *Vault) readPrivateKey(path string) error {
	pemData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		return errors.New("Unable to decode private key")
	}
	if v.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return err
	}
	return nil
}
