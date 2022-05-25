package encrypt

import (
	"encoding/base64"
	"fmt"
	"github.com/seanbit/gokit/fileutils"
	"log"
	"testing"
)

func TestRsaEncryptImpl_Encrypt(t *testing.T) {
	fileutils.CheckExist("")
	buf, err := fileutils.ReadFile("/Users/lyra/Desktop/Doc/安全方案/businessS/spubkey.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(buf))

	src := "this is a test word for rsa encrypt"
	encryptData, err := Rsa().Encrypt(string(buf), []byte(src))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encryptData))
}

func TestRsaEncryptImpl_Decrypt(t *testing.T) {

	buf, err := fileutils.ReadFile("/Users/lyra/Desktop/Doc/安全方案/businessS/sprivkey.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(buf))

	src := "SWD8oHiSHA645ez1isB8fuXy6JLhDgQfDbvWHUUYDswg0qeTV6i3g9dQ/yZBMd0UbjEpmo03D9dS54WMAF4BGVRtkizJiecqxL4Hm6O4hWqSzaQxunIcv2seC5qmJbVLP4SNvv+Y/BQ9k5me9mqS7W0xucb3Jj6U2FqDybU2+9E="
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		t.Error(err)
	}
	decryptData, err := Rsa().Decrypt(string(buf), data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(decryptData))
}

func TestRsaEncryptImpl_Sign(t *testing.T) {

	buf, err := fileutils.ReadFile("/Users/lyra/Desktop/Doc/安全方案/businessS/sprivkey.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(buf))

	src := "this is a test word for rsa sign"
	signData, err := Rsa().Sign(string(buf), []byte(src))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(signData))
}

func TestRsaEncryptImpl_Verify(t *testing.T) {

	buf, err := fileutils.ReadFile("/Users/lyra/Desktop/Doc/安全方案/businessS/spubkey.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(buf))

	src := "this is a test word for rsa sign"
	signStr := "Sdg95tT3TRYZu21wwm6rPup3dnh4eGz77iK9ikG9lLuWp1Obbi4L5SEnRuE/aqC23b1lNVDheOTHQuTZhTkoY5NzWSvOv05h3+MmXYHgNO+329X16K/vAl5RagEGV6P9dNrJH1yjX95MzNfXTguG9ZCRDC6kthYrzLk+3Z7mr5w="
	signData, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		t.Error(err)
	}
	err = Rsa().Verify(string(buf), []byte(src), signData)
	if err != nil {
		t.Error(err)
	} else {
		log.Println("verify success")
	}
}

func TestRsaEncryptImpl_Verify2(t *testing.T) {

	buf, err := fileutils.ReadFile("/Users/lyra/Desktop/Doc/安全方案/businessC/cpubkey.pem")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(buf))

	str := "haha"
	src := []byte(str)
	signStr := "vVl0WXegq4xfpAiDFIMLulhfc2lxaiewSdZ9gJbJGYqPYqmvQkzxJXbSrWxtd+YRjjrGe2+KShENYDP74vgplbnF/fDdG2XnSdQW/VltlRTQu2Z5VOzYlJgDBIbuynYlGIk3Qnc74TLmoi2qd/fWPvFAMHtaccR6EgUvjeeuO54="
	signData, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		t.Error(err)
	}
	err = Rsa().Verify(string(buf), src, signData)
	if err != nil {
		t.Error(err)
	} else {
		log.Println("verify success")
	}
}