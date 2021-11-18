package test

import (
	"testing"
	"github.com/simonalong/go-util"
)

func TestMapToProperties1(t *testing.T) {

	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = 13
	dataMap["c"] = 14

	go-util.MapToProperties(dataMap)
}
