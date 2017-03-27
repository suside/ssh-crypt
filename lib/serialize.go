package lib

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/gob"
	"hash/crc32"
)

func init() {
	gob.Register([]*rsa.PublicKey{})
	gob.Register(rsa.PublicKey{})
	gob.Register(rsa.PrivateKey{})
	gob.Register(vaultSecured{})
}

func toBytes(v interface{}) []byte {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(&v)
	return b.Bytes()
}

func toBase64(v interface{}) string {
	return base64.StdEncoding.EncodeToString(toBytes(v))
}

func fromBase64(str string) (interface{}, error) {
	var m interface{}
	by, _ := base64.StdEncoding.DecodeString(str)
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	if err := d.Decode(&m); err != nil {
		return m, err
	}
	return m, nil
}

func crc32sum(v interface{}) uint32 {
	return crc32.Checksum(toBytes(v), crc32.IEEETable)
}
