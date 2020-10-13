package foundation

import (
	"fmt"
	"testing"
)

type p struct {
	string
}

func (this *p) print() {
	fmt.Println(this.string)
}
func TestRG(t *testing.T) {
	var service = &p{"service"}
	Register(iexType, service)
	s := Get(iexType).(iex)
	s.print()
}