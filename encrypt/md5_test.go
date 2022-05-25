package encrypt

import (
	"fmt"
	"testing"
)

func TestMd5EncryptImpl_EncryptWithTimestamp(t *testing.T) {
	key := "ajsdhjbzjxcbzhcb"
	value := "this is test value"
	enc := Md5().Encode([]byte(value))
	fmt.Println(enc)
	enc = Md5().HmacEncode([]byte(key), []byte(value))
	fmt.Println(enc)
}