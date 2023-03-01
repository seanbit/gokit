package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

/**
 * You can generate binary exec pkg Use: >go test -c pem_test.go pem.go utils_idea.go -o pem.test
 * For TestGenRsaKeyPairForShow 	Run: >./pem.test -test.run "TestGenRsaKeyPairForShow" -test.v
 * For TestGenRsaKeyPairForExport 	Run: >./pem.test -test.run "TestGenRsaKeyPairForExport" -test.v
 * For TestGenRsaKeyPairForServer 	Run: >./pem.test -test.run "TestGenRsaKeyPairForServer" -test.v
 */

func TestGenRsaKeyPairForShow(t *testing.T) {
	pem, key, err := GenRsaKey()
	if err != nil {
		t.Error(err)
		return
	}
	if err := testRsaKeyPairValidate(string(pem), string(key)); err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(pem))
	fmt.Println(string(key))
}

func TestGenRsaKeyPairForExport(t *testing.T) {
	pem, key, err := GenRsaKey()
	if err != nil {
		t.Error(err)
		return
	}
	if err := testRsaKeyPairValidate(string(pem), string(key)); err != nil {
		t.Fatal(err)
	}
	// export
	if err := ioutil.WriteFile("./pub.pem", pem, fs.ModePerm); err != nil {
		t.Fatal(err)
	} else {
		t.Log("pem export success")
	}
	if err := ioutil.WriteFile("./key.pem", key, fs.ModePerm); err != nil {
		t.Fatal(err)
	} else {
		t.Log("key export success")
	}
}

func TestGenRsaKeyPairForServer(t *testing.T) {
	pem, key, err := GenRsaKey()
	if err != nil {
		t.Fatal(err)
	}
	if err := testRsaKeyPairValidate(string(pem), string(key)); err != nil {
		t.Fatal(err)
	}
	// idea key
	k := GenerateIdeaKey()
	t.Log(hex.EncodeToString(k))
	encKey, err := KeyEncrypt(k, key)
	if err != nil {
		t.Fatal(err)
	}
	decKey, err := KeyDecrypt(k, encKey)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(key, decKey) {
		t.Fatal("idea dec key not right, please run again.")
	}
	// export
	filePathName := "./server"
	if err := ioutil.WriteFile(filePathName+".pem", pem, fs.ModePerm); err != nil {
		t.Fatal(err)
	} else {
		t.Log("server.pem export success")
	}
	if err := ioutil.WriteFile(filePathName+".key", []byte(base64.StdEncoding.EncodeToString(encKey)), fs.ModePerm); err != nil {
		t.Fatal(err)
	} else {
		t.Log("server.key export success")
	}
}
