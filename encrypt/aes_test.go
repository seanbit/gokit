package encrypt

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestAesEncryptImpl_EncryptCBC(t *testing.T) {

	key := Aes().GenerateKey()
	fmt.Println(key)
	fmt.Println(hex.EncodeToString(key))
	fmt.Println(base64.StdEncoding.EncodeToString(key))

	origData := []byte("Hello World") // 待加密的数据

	log.Println("------------------ CBC模式 --------------------")
	encrypted, err := Aes().EncryptCBC(origData, key)
	if err != nil {
		t.Error(err)
	}
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))

	decrypted, err := Aes().DecryptCBC(encrypted, key)
	if err != nil {
		t.Error(err)
	}
	log.Println("解密结果：", string(decrypted))

	//log.Println("------------------ ECB模式 --------------------")
	//encrypted = AesEncryptECB(origData, key)
	//log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	//log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	//decrypted = AesDecryptECB(encrypted, key)
	//log.Println("解密结果：", string(decrypted))
	//
	//log.Println("------------------ CFB模式 --------------------")
	//encrypted = AesEncryptCFB(origData, key)
	//log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	//log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	//decrypted = AesDecryptCFB(encrypted, key)
	//log.Println("解密结果：", string(decrypted))

}
