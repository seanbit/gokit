package utils

import (
	"crypto/rc4"
	"fmt"
	"math/big"
	"math/rand"
	"net/url"
	"testing"
	"time"
)

func TestNumToBHex(t *testing.T) {
	//var num uint64 = uint64(668801) + uint64(12345678901001) + uint64(time.Now().Unix())
	//fmt.Println(num)
	fmt.Println(NumsToBHex([]uint64{668849, 236644620316971008}))
	fmt.Println(NumsToBHex([]uint64{12345678901001, 12345678901001}))
}

func TestBHex2Num(t *testing.T) {
	fmt.Println(BHex2Nums("FKYYZ2FRHRTCRK4VN"))
	fmt.Println(BHex2Nums("5GUWTNKAQZ5GUWTNKAQ"))
}

func TestRandString(t *testing.T) {
	for idx := 0; idx < 10; idx++ {
		fmt.Println(RandString(12))
	}
}

func TestRandSmsCode(t *testing.T) {
	c := NewValidateCode(6)
	fmt.Println(c)
}

func TestParseQuery(t *testing.T) {
	q := "requestid=kzhnlakr!@#I(#-12a(!@U&userid=1001&username=0xzhangsan"
	vals, err := url.ParseQuery(q)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(vals)
	t.Log(vals.Get("requestid"))
	t.Log(vals.Get("userid"))
	t.Log(vals.Get("username"))
}

func TestUtsMinute(t *testing.T) {
	for idx := 0; idx < 10; idx++ {
		t.Log(UtsMinute())
		time.Sleep(time.Minute)
	}
}

func TestUtsMinute2(t *testing.T) {
	tm := time.Date(2158, 12, 12, 23, 59, 58, 698, time.Local)
	second := tm.Unix()
	t.Log(second / 60)
}

func TestRC4(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	for idx := 0; idx < 1000000; idx++ {

		// 密钥
		key := big.NewInt(int64(rand.Uint64()))
		_, err := rc4.NewCipher(key.Bytes())
		if err != nil {
			t.Fatalf("key error : %d", key)
		}

		// 要加密的源数据
		src := big.NewInt(int64(rand.Uint64()))
		dst := Rc4Encrypt(key, src)
		ori, err := Rc4Decrypt(key, dst)
		if err != nil {
			t.Fatalf("decrypt error : %s", err)
		}

		t.Logf("key: %d, src: %d, dst: %s, ori: %d\n", key.Uint64(), src.Uint64(), dst, ori.Uint64())

		if ori.Uint64() != src.Uint64() {
			t.Fatalf("src: %d not equal to ori: %d\n", src.Uint64(), ori.Uint64())
		}
	}
}

func TestMyRand(t *testing.T) {
	var rd = NewRand(16888, time.Second*30)
	for idx := 0; idx < 100; idx++ {
		t.Log(rd.RandInt(10))
	}
}

func TestGetDay(t *testing.T) {
	t.Log(GetDay(time.Now()))
}

type TestEncParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestRandStr(t *testing.T) {
	fmt.Println(RandStringBytes(64))
}
