package test

import (
	"fmt"
	"github.com/simonalong/gole/util"
	"reflect"
	"testing"
)

func TestField(t *testing.T) {
	privateField := PrivateFieldStruct{}
	data := "data"
	util.SetFieldPrivateValue(reflect.ValueOf(privateField), "name", reflect.ValueOf(&data))

	dataRel := util.GetPrivateFieldValue(reflect.ValueOf(&privateField), "name")
	fmt.Println(dataRel)
}

type PrivateFieldStruct struct {
	name string
	age  int
}
