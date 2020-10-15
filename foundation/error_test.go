package foundation

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseError(t *testing.T) {
	err := NewError(errors.New("haha-a, asd@!()uw-  wow! 这是汉语哎，還有繁體字呢。 hey834w"), 101, "some wrong")
	e, ok := ParseError(err)
	if ok {
		fmt.Printf("code:%d, msg:%s, err:%s", e.Code(), e.Msg(), e.Error())
	} else {
		fmt.Println("not Error")
	}
}

func TestParseError2(t *testing.T) {
	err := NewError(errors.New("haha-a, asd@!()uw-  wow! 这是汉语哎，還有繁體字呢。 hey834w"), 101, "")
	e, ok := ParseError(err)
	if ok {
		fmt.Printf("code:%d, msg:%s, err:%s", e.Code(), e.Msg(), e.Error())
	} else {
		fmt.Println("not Error")
	}
}

func TestParseError3(t *testing.T) {
	err := NewError(nil, 101, "some wrong")
	e, ok := ParseError(err)
	if ok {
		fmt.Printf("code:%d, msg:%s, err:%s", e.Code(), e.Msg(), e.Error())
	} else {
		fmt.Println("not Error")
	}
}