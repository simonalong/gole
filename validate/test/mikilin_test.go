package main

import (
	"github.com/simonalong/gole/validate"
	"testing"
)

type ValidateEntity struct {
	Name string `match:"value={zhou, 宋江} isBlank=true"`
	Age  int    `match:"value={12, 13}"`
}

type ValidateEntity2 struct {
	Name string `match:"value={zhou, 宋江} isBlank"`
	Age  int    `match:"value={12, 13}"`
}

func TestValidateGole1(t *testing.T) {
	var value ValidateEntity
	var result bool
	var err string

	//测试 正常情况
	value = ValidateEntity{Age: 12}
	result, _, err = validate.Check(value)
	TrueErr(t, result, err)

	// 测试 正常情况
	value = ValidateEntity{Age: 13, Name: "zhou"}
	result, _, err = validate.Check(value)
	TrueErr(t, result, err)

	// 测试 异常情况
	value = ValidateEntity{Age: 13, Name: "陈真"}
	result, _, err = validate.Check(value)
	Equal(t, err, "[\"属性 Name 的值 陈真 不在只可用列表 [zhou 宋江] 中\",\"属性 Name 的值为非空字符\"]", result, false)
}

func TestValidateGole2(t *testing.T) {
	var value ValidateEntity2
	var result bool
	var err string

	//测试 正常情况
	value = ValidateEntity2{Age: 12}
	result, _, err = validate.Check(value)
	TrueErr(t, result, err)

	// 测试 正常情况
	value = ValidateEntity2{Age: 13, Name: "zhou"}
	result, _, err = validate.Check(value)
	TrueErr(t, result, err)

	// 测试 异常情况
	value = ValidateEntity2{Age: 13, Name: "陈真"}
	result, _, err = validate.Check(value)
	Equal(t, "[\"属性 Name 的值 陈真 不在只可用列表 [zhou 宋江] 中\",\"属性 Name 的值为非空字符\"]", err, result, false)
}

// 压测进行基准测试
func BenchmarkValidateGole3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validate.Check(ValidateEntity{Age: 12})
	}
}
