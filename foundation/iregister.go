package foundation

import (
	"reflect"
)

var itypeObjMap = make(map[reflect.Type]interface{})

func Register(itype reflect.Type, obj interface{}) {
	if itype == nil || obj == nil {
		return
	}
	itypeObjMap[itype] = obj
}

func Get(interType reflect.Type) (obj interface{}) {
	return itypeObjMap[interType]
}

type iex interface {
	print()
}
var iexType = reflect.TypeOf((*iex)(nil)).Elem()

