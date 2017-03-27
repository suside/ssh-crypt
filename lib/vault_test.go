package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readAuthorizedKeys(t *testing.T) {
	v := Vault{}
	v.ReadAuthorizedKeys("../test/authorized_keys")
	assert.Equal(t, uint32(0x752098b2), crc32sum(v.publicKeys[0]))
	assert.Equal(t, uint32(0xbc02a9b3), crc32sum(v.publicKeys[1]))
}

func Test_readPrivateKey(t *testing.T) {
	v := Vault{}
	v.readPrivateKey("../test/t1_id_rsa")
	assert.Equal(t, uint32(0xaa0afb16), crc32sum(v.privateKey))
}

func Test_storeAndReadVault(t *testing.T) {
	tmpfile, _ := ioutil.TempFile("", "ssh-crypt-test")
	defer os.Remove(tmpfile.Name())
	v1 := Vault{Plaintext: []byte("my secret")}
	v1.ReadAuthorizedKeys("../test/authorized_keys")
	v1.StoreSecuredVault(tmpfile.Name())

	v2 := Vault{}
	v2.DecryptVaultWithKey(tmpfile.Name(), "../test/t1_id_rsa")
	v3 := Vault{}
	v3.DecryptVaultWithKey(tmpfile.Name(), "../test/t2_id_rsa")
	assert.Equal(t, v2.Plaintext, v3.Plaintext)
}

func Test_storeAndReadVaultWithNotUsedKey(t *testing.T) {
	tmpfile, _ := ioutil.TempFile("", "ssh-crypt-test")
	defer os.Remove(tmpfile.Name())
	v1 := Vault{Plaintext: []byte("my secret")}
	v1.ReadAuthorizedKeys("../test/authorized_keys")
	v1.StoreSecuredVault(tmpfile.Name())

	v3 := Vault{}
	err := v3.DecryptVaultWithKey(tmpfile.Name(), "../test/t3_id_rsa")
	assert.Equal(t, "Unable to read vault "+tmpfile.Name()+" with key ../test/t3_id_rsa", err.Error())
}

func Test_EncodingVaultSecured(t *testing.T) {
	v := vaultSecured{}
	v1, err := fromBase64(toBase64(v))
	assert.Equal(t, v, v1)
	assert.Nil(t, err)
}

func Test_ReadWithKeyEmptyVault(t *testing.T) {
	v := Vault{}
	v.DecryptVaultWithKey("/tmp/newvault123123", "../test/t1_id_rsa")
	assert.Equal(t, []byte(""), v.Plaintext)
}

func Test_ReadWithKeyNotBase64(t *testing.T) {
	v := Vault{}
	assert.Equal(
		t,
		"../test/authorized_keys content does not look like base64",
		v.DecryptVaultWithKey("../test/authorized_keys", "../test/t1_id_rsa").Error(),
	)
}

func Test_ReadWithKeyNoKey(t *testing.T) {
	v := Vault{}
	assert.Equal(
		t,
		"no such file",
		v.DecryptVaultWithKey("../test/authorized_keys", "no such file").(*os.PathError).Path,
	)
}

func Test_readPrivateKeyNoKey(t *testing.T) {
	v := Vault{}
	assert.Equal(t, "no such file", v.readPrivateKey("no such file").(*os.PathError).Path)
}

func Test_readPrivateKeyFail2(t *testing.T) {
	v := Vault{}
	assert.Equal(t, errors.New("Unable to decode private key"), v.readPrivateKey("../test/authorized_keys"))
}
