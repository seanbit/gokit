package foundation

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

type Error interface {
	error
	Code() int
	Msg() string
}

type cerr struct {
	error
	code 	int
	msg  	string
}

func (this *cerr) Code() int {
	return this.code
}

func (this *cerr) Msg() string {
	return this.msg
}

func (this *cerr) Unwrap() error {
	return this.error
}

func NewError(err error, code int, msg string) error {
	if err == nil {
		err = errors.New("")
	}
	return fmt.Errorf("code:%d, msg:%s, err:%s", code, msg, err.Error())
}

func ParseError(err error) (Error, bool) {
	if err == nil {
		return nil, false
	}
	errRg := regexp.MustCompile(`^code:(\d+), msg:([\s\S]+), err:([\s\S]+)`)
	params := errRg.FindStringSubmatch(err.Error())
	if params[0] != err.Error() {
		return nil, false
	}
	if len(params) != 4 {
		return nil, false
	}
	code, err := strconv.Atoi(params[1])
	if err != nil {
		return nil, false
	}
	return &cerr{
		error: errors.New(params[3]),
		code:  code,
		msg:   params[2],
	}, true
}

func TestParseError(t *testing.T) {
	err := NewError(errors.New("haha-a, asd@!()uw-  wow! 这是汉语哎，還有繁體字呢。 hey834w"), 101, "some wrong")
	e, ok := ParseError(err)
	if ok {
		fmt.Printf("code:%d, msg:%s, err:%s", e.Code(), e.Msg(), e.Error())
	} else {
		fmt.Println("not Error")
	}
}
