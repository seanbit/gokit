package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/seanbit/gokit/encrypt"
	"time"
)

const (
	keybits = 2048
	days    = time.Duration(1024)
)

//RSA公钥私钥产生
func GenRsaKey() (pub, key []byte, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, keybits)
	if err != nil {
		return nil, nil, err
	}
	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, nil, err
	}

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	//file, err := os.Create("private.pem")
	//if err != nil {
	//	return err
	//}
	keyFile := new(bytes.Buffer)
	err = pem.Encode(keyFile, block)
	if err != nil {
		return nil, nil, err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubFile := new(bytes.Buffer)
	err = pem.Encode(pubFile, block)
	if err != nil {
		return nil, nil, err
	}
	return pubFile.Bytes(), keyFile.Bytes(), nil
}

func GenerateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 32)
	copy(genKey, key)
	for i := 32; i < len(key); {
		for j := 0; j < 32 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func KeyEncrypt(k, key []byte) ([]byte, error) {
	const encLen = 8
	encRounds := len(key) / encLen
	encKey := make([]byte, len(key))
	for round := 0; round < encRounds; round++ {
		cipher, err := NewCipher(k)
		if err != nil {
			return nil, err
		}
		cipher.Encrypt(encKey[round*8:(round+1)*8], key[round*8:(round+1)*8])
	}
	return encKey, nil
}

func KeyDecrypt(k, key []byte) ([]byte, error) {
	const encLen = 8
	decRounds := len(key) / encLen
	decKey := make([]byte, len(key))
	for round := 0; round < decRounds; round++ {
		cipher, err := NewCipher(k)
		if err != nil {
			return nil, err
		}
		cipher.Decrypt(decKey[round*8:(round+1)*8], key[round*8:(round+1)*8])
	}
	return decKey, nil
}

func testRsaKeyPairValidate(pub, key string) error {
	var params = map[string]string{
		"k1": "v1",
	}
	jsonBts, _ := json.Marshal(params)
	enBts, err := encrypt.Rsa().Encrypt(pub, jsonBts)
	if err != nil {
		return err
	}

	rawBts, err := encrypt.Rsa().Decrypt(key, enBts)
	if err != nil {
		return err
	}
	if !bytes.Equal(jsonBts, rawBts) {
		return errors.New("json bts not equal raw bts")
	}

	sign, err := encrypt.Rsa().Sign(key, jsonBts)
	if err != nil {
		return err
	}
	err = encrypt.Rsa().Verify(pub, jsonBts, sign)
	if err != nil {
		return err
	}
	return nil
}
