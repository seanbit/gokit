package utils

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var pwdSrc = rand.NewSource(time.Now().UnixNano())

func RandString(n int) string {
	b := make([]byte, n)
	// A pwdSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, pwdSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = pwdSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func UtsMinute() int64 {
	second := time.Now().Unix()
	return second / 60
}

func UtsDay() int64 {
	second := time.Now().Unix()
	return second / 60 / 60 / 24
}

func getMonth() uint32 {
	t := time.Now()
	ts := fmt.Sprintf("%04d%02d", t.Year(), t.Month())
	month, _ := strconv.ParseUint(ts, 10, 64)
	return uint32(month)
}

func getWeek() uint32 {
	year, week := time.Now().ISOWeek()
	ts := fmt.Sprintf("%04d%02d", year, week)
	w, _ := strconv.ParseUint(ts, 10, 64)
	return uint32(w)
}

func GetDay(tm time.Time) uint32 {
	return uint32(tm.Year()*10000 + int(tm.Month())*100 + tm.Day())
}

func ParseDay(formatDay uint32) time.Time {
	if formatDay < 20000101 {
		return time.Time{}
	}
	year := formatDay / 10000
	monthDay := formatDay % 10000
	month := monthDay / 100
	day := monthDay % 100
	local := time.Now().Location()
	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, local)
}

func Rc4Encrypt(key, src *big.Int) string {
	// 加密操作
	dstBts := make([]byte, len(src.Bytes()))
	cipher, _ := rc4.NewCipher(key.Bytes())
	cipher.XORKeyStream(dstBts, src.Bytes())
	return hex.EncodeToString(dstBts)
}

func Rc4Decrypt(key *big.Int, hexSrc string) (*big.Int, error) {
	// 解密操作
	srcBts, err := hex.DecodeString(hexSrc)
	if err != nil {
		return nil, err
	}
	dstBts := make([]byte, len(srcBts))
	cipher, _ := rc4.NewCipher(key.Bytes()) // 切记：这里不能重用cipher1，必须重新生成新的
	cipher.XORKeyStream(dstBts, srcBts)
	dst := big.NewInt(0)
	dst.SetBytes(dstBts)
	return dst, nil
}
